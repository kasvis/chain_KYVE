package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SubmitBundleProposal handles the logic of an SDK message that allows protocol nodes to submit a new bundle proposal.
func (k msgServer) SubmitBundleProposal(
	goCtx context.Context, msg *types.MsgSubmitBundleProposal,
) (*types.MsgSubmitBundleProposalResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, found := k.GetPool(ctx, msg.Id)

	// Error if the pool isn't found.
	if !found {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), msg.Id)
	}

	// Check if enough nodes are online
	if len(pool.Stakers) < 2 {
		return nil, types.ErrNotEnoughNodesOnline
	}

	// Error if the pool has no funds.
	if len(pool.Funders) == 0 {
		return nil, sdkErrors.Wrap(sdkErrors.ErrInsufficientFunds, types.ErrFundsTooLow.Error())
	}

	// Error if the pool is paused.
	if pool.Paused {
		return nil, sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrPoolPaused.Error())
	}

	// Check if the sender is a protocol node (aka has staked into this pool).
	_, isStaker := k.GetStaker(ctx, msg.Creator, msg.Id)
	if !isStaker {
		return nil, sdkErrors.Wrap(sdkErrors.ErrUnauthorized, types.ErrNoStaker.Error())
	}

	// Validate bundle id.
	if msg.BundleId == "" {
		return nil, types.ErrInvalidArgs
	}

	// resubmit NO_DATA_BUNDLE
	if pool.BundleProposal.BundleId == types.NO_DATA_BUNDLE && pool.BundleProposal.Uploader == msg.Creator {
		// Check if bundle id is an ARWEAVE_BUNDLE
		if msg.BundleId == types.NO_QUORUM_BUNDLE || msg.BundleId == types.NO_DATA_BUNDLE  {
			return nil, types.ErrInvalidArgs
		}

		// Validate bundle args
		if msg.BundleSize == 0 || msg.ByteSize == 0 {
			return nil, types.ErrInvalidArgs
		}

		// resubmit ARWEAVE_BUNDLE
		pool.BundleProposal = &types.BundleProposal{
			Uploader:     pool.BundleProposal.Uploader,
			NextUploader: pool.BundleProposal.NextUploader,
			BundleId:     msg.BundleId,
			ByteSize:     msg.ByteSize,
			FromHeight:   pool.BundleProposal.ToHeight,
			ToHeight:     pool.BundleProposal.ToHeight + msg.BundleSize,
			CreatedAt:    uint64(ctx.BlockTime().Unix()),
		}

		k.SetPool(ctx, pool)

		return &types.MsgSubmitBundleProposalResponse{}, nil
	}

	// Check if upload_interval has been surpassed
	if uint64(ctx.BlockTime().Unix()) < (pool.BundleProposal.CreatedAt + pool.UploadInterval) {
		return nil, types.ErrUploadInterval
	}

	// Check if the sender is the designated uploader.
	if pool.BundleProposal.NextUploader != msg.Creator {
		return nil, types.ErrNotDesignatedUploader
	}

	// Check if consensus has already been reached.
	valid := false
	invalid := false

	if len(pool.Stakers) > 1 {
		// subtract one because of uploader
		valid = len(pool.BundleProposal.VotersValid)*2 > (len(pool.Stakers) - 1)
		invalid = len(pool.BundleProposal.VotersInvalid)*2 >= (len(pool.Stakers) - 1)
	}

	// If the next-uploader submits the NO_QUORUM_BUNDLE it means that
	// there was no quorum reached in the current round.
	if msg.BundleId == types.NO_QUORUM_BUNDLE {
		// Validate bundle args
		if msg.BundleSize != 0 || msg.ByteSize != 0 {
			return nil, types.ErrInvalidArgs
		}

		// check if the quorum was actually reached
		if valid || invalid {
			// next_uploader is submitting an invalid bundle
			// because the quorum was reached
			// TODO add some tolerance -> maybe in pool_parameters

			// slash next_uploader for false NO_QUORUM_BUNDLE submit
			slashAmount := k.slashStaker(ctx, &pool, msg.Creator, k.UploadSlash(ctx))

			// emit slashing event
			types.EmitSlashEvent(ctx, pool.Id, msg.Creator, slashAmount)

			// Update current lowest staker
			k.updateLowestStaker(ctx, &pool)

			// keep the BundleProposal
			// just replace the nextUploader and the createdAt time.
			pool.BundleProposal.CreatedAt = uint64(ctx.BlockTime().Unix())
			// select next_uploader from voters and uploader
			candidates := append(append(pool.BundleProposal.VotersValid, pool.BundleProposal.VotersInvalid...), pool.BundleProposal.Uploader)
			pool.BundleProposal.NextUploader = k.getNextUploaderByRandom(ctx, &pool, candidates)
		} else {
			// If consensus wasn't reached, we drop the bundle and emit an event.
			types.EmitBundleDroppedQuorumNotReachedEvent(ctx, &pool)

			pool.BundleProposal = &types.BundleProposal{
				NextUploader: k.getNextUploaderByRandom(ctx, &pool, pool.Stakers),
				FromHeight:   pool.BundleProposal.FromHeight,
				ToHeight:  pool.BundleProposal.FromHeight,
				CreatedAt: uint64(ctx.BlockTime().Unix()),
			}
		}

		k.SetPool(ctx, pool)

		return &types.MsgSubmitBundleProposalResponse{}, nil
	}

	// Check args of NO_DATA_BUNDLE
	if msg.BundleId == types.NO_DATA_BUNDLE {
		// Validate bundle args
		if msg.BundleSize != 0 || msg.ByteSize != 0 {
			return nil, types.ErrInvalidArgs
		}
	}

	// If the pool is in "genesis state" or the bundle was dropped, just register new proposal.
	if pool.BundleProposal.BundleId == "" {
		pool.BundleProposal = &types.BundleProposal{
			Uploader:     msg.Creator,
			NextUploader: k.getNextUploaderByRandom(ctx, &pool, pool.Stakers),
			BundleId:     msg.BundleId,
			ByteSize:     msg.ByteSize,
			FromHeight:   pool.BundleProposal.ToHeight,
			ToHeight:     pool.BundleProposal.ToHeight + msg.BundleSize,
			CreatedAt:    uint64(ctx.BlockTime().Unix()),
		}

		k.SetPool(ctx, pool)

		return &types.MsgSubmitBundleProposalResponse{}, nil
	}

	// EVALUATE PREVIOUS ROUND

	// handle valid proposal
	if valid {
		// Calculate the total reward for the bundle, and individual payouts.
		bundleReward := pool.OperatingCost + (pool.BundleProposal.ByteSize * k.StorageCost(ctx))

		// load and parse network fee
		networkFee, err := sdk.NewDecFromStr(k.NetworkFee(ctx))
		if err != nil {
			k.PanicHalt(ctx, "Invalid value for params: "+err.Error())
		}

		treasuryPayout := uint64(sdk.NewDec(int64(bundleReward)).Mul(networkFee).RoundInt64())
		uploaderPayout := bundleReward - treasuryPayout

		// Calculate the delegation rewards for the uploader.
		uploader, foundUploader := k.GetStaker(ctx, pool.BundleProposal.Uploader, pool.Id)
		uploaderDelegation, foundUploaderDelegation := k.GetDelegationPoolData(ctx, pool.Id, pool.BundleProposal.Uploader)

		if foundUploader && foundUploaderDelegation {
			// If the uploader has no delegators, it keeps the delegation reward.

			if uploaderDelegation.DelegatorCount > 0 {
				// Calculate the reward, factoring in the node commission, and subtract from the uploader payout.
				commission, _ := sdk.NewDecFromStr(uploader.Commission)
				delegationReward := uint64(
					sdk.NewDec(int64(uploaderPayout)).Mul(sdk.NewDec(1).Sub(commission)).RoundInt64(),
				)

				uploaderPayout -= delegationReward
				uploaderDelegation.CurrentRewards += delegationReward

				k.SetDelegationPoolData(ctx, uploaderDelegation)
			}
		}

		// Calculate the individual cost for each pool funder.
		// NOTE: Because of integer division, it is possible that there is a small remainder.
		//       This remainder is in worst case MaxFundersAmount(tkyve) and is charged to the lowest funder.
		fundersCost := bundleReward / uint64(len(pool.Funders))
		fundersCostRemainder := bundleReward - uint64(len(pool.Funders))*fundersCost

		// Fetch the lowest funder, and find a new one if the current one isn't found.
		lowestFunder, foundLowestFunder := k.GetFunder(ctx, pool.LowestFunder, pool.Id)

		if !foundLowestFunder {
			k.updateLowestFunder(ctx, &pool)
			lowestFunder, _ = k.GetFunder(ctx, pool.LowestFunder, pool.Id)
		}

		// Remove every funder who can't afford the funder cost.
		if fundersCost+fundersCostRemainder > lowestFunder.Amount {
			// First, let's remove the lowest funder.
			k.removeFunder(ctx, &pool, &lowestFunder)

			err := k.transferToTreasury(ctx, lowestFunder.Amount)
			if err != nil {
				return nil, err
			}

			// Now, let's remove all other funders who have run out of funds.
			for _, account := range pool.Funders {
				funder, _ := k.GetFunder(ctx, account, pool.Id)

				if funder.Amount < fundersCost {
					k.removeFunder(ctx, &pool, &funder)

					err := k.transferToTreasury(ctx, funder.Amount)
					if err != nil {
						return nil, err
					}
				}
			}

			// Recalculate the lowest funder, update, and return.
			k.updateLowestFunder(ctx, &pool)

			pool.BundleProposal = &types.BundleProposal{
				Uploader:      pool.BundleProposal.Uploader,
				NextUploader:  pool.BundleProposal.NextUploader,
				BundleId:      pool.BundleProposal.BundleId,
				ByteSize:      pool.BundleProposal.ByteSize,
				FromHeight:    pool.BundleProposal.FromHeight,
				ToHeight:      pool.BundleProposal.ToHeight,
				CreatedAt:     uint64(ctx.BlockTime().Unix()),
				VotersValid:   pool.BundleProposal.VotersValid,
				VotersInvalid: pool.BundleProposal.VotersInvalid,
			}

			k.SetPool(ctx, pool)

			// Emit a bundle dropped event because of insufficient funds.
			types.EmitBundleDroppedInsufficientFundsEvent(ctx, &pool)

			return &types.MsgSubmitBundleProposalResponse{}, nil
		}

		// Charge every funder equally.
		for _, account := range pool.Funders {
			funder, _ := k.GetFunder(ctx, account, pool.Id)
			funder.Amount -= fundersCost
			k.SetFunder(ctx, funder)
		}

		// Remove any remainder cost from the lowest funder.
		lowestFunder, _ = k.GetFunder(ctx, pool.LowestFunder, pool.Id)
		lowestFunder.Amount -= fundersCostRemainder
		k.SetFunder(ctx, lowestFunder)

		// Subtract bundle reward from the pool's total funds.
		pool.TotalFunds -= bundleReward

		// Only slash with vote slash if bundle is of type ARWEAVE_BUNDLE
		if pool.BundleProposal.BundleId != types.NO_DATA_BUNDLE {
			// Partially slash all nodes who voted incorrectly.
			for _, voter := range pool.BundleProposal.VotersInvalid {
				slashAmount := k.slashStaker(ctx, &pool, voter, k.VoteSlash(ctx))
				types.EmitSlashEvent(ctx, pool.Id, voter, slashAmount)
			}
		}

		// Send payout to treasury.
		errTreasury := k.transferToTreasury(ctx, treasuryPayout)
		if errTreasury != nil {
			return nil, errTreasury
		}

		// Send payout to uploader.
		errTransfer := k.TransferToAddress(ctx, pool.BundleProposal.Uploader, uploaderPayout)
		if errTransfer != nil {
			return nil, errTransfer
		}

		// Finalise the proposal, saving useful information.
		pool.HeightArchived = pool.BundleProposal.ToHeight
		pool.BytesArchived = pool.BytesArchived + pool.BundleProposal.ByteSize
		pool.TotalBundleRewards = pool.TotalBundleRewards + bundleReward
		pool.TotalBundles = pool.TotalBundles + 1

		// Only add bundles with data to store
		if pool.BundleProposal.BundleId != types.NO_DATA_BUNDLE {
			k.SetProposal(ctx, types.Proposal{
				BundleId:    pool.BundleProposal.BundleId,
				PoolId:      pool.Id,
				Uploader:    pool.BundleProposal.Uploader,
				FromHeight:  pool.BundleProposal.FromHeight,
				ToHeight:    pool.BundleProposal.ToHeight,
				FinalizedAt: uint64(ctx.BlockHeight()),
			})
		}

		// Emit a valid bundle event.
		types.EmitBundleValidEvent(ctx, &pool, bundleReward)

		// select next_uploader from voters and uploader
		candidates := append(append(pool.BundleProposal.VotersValid, pool.BundleProposal.VotersInvalid...), pool.BundleProposal.Uploader)

		// Set submitted bundle as new bundle proposal and select new next_uploader
		pool.BundleProposal = &types.BundleProposal{
			Uploader:     msg.Creator,
			NextUploader: k.getNextUploaderByRandom(ctx, &pool, candidates),
			BundleId:     msg.BundleId,
			ByteSize:     msg.ByteSize,
			FromHeight:   pool.BundleProposal.ToHeight,
			ToHeight:     pool.BundleProposal.ToHeight + msg.BundleSize,
			CreatedAt:    uint64(ctx.BlockTime().Unix()),
		}

		k.SetPool(ctx, pool)

		return &types.MsgSubmitBundleProposalResponse{}, nil
	} else if invalid {
		// Only slash with vote slash if bundle is of type ARWEAVE_BUNDLE
		if pool.BundleProposal.BundleId != types.NO_DATA_BUNDLE {
			// Partially slash all nodes who voted incorrectly.
			for _, voter := range pool.BundleProposal.VotersValid {
				slashAmount := k.slashStaker(ctx, &pool, voter, k.VoteSlash(ctx))
				types.EmitSlashEvent(ctx, pool.Id, voter, slashAmount)
			}
		}

		var slashAmount uint64

		// Partially slash the uploader.
		if pool.BundleProposal.BundleId == types.NO_DATA_BUNDLE {
			// Slash uploader with timeout slash if bundle was of type NO_DATA_BUNDLE
			slashAmount = k.slashStaker(ctx, &pool, pool.BundleProposal.Uploader, k.TimeoutSlash(ctx))
		} else {
			// Slash uploader with upload slash if bundle was of type ARWEAVE_BUNDLE
			slashAmount = k.slashStaker(ctx, &pool, pool.BundleProposal.Uploader, k.UploadSlash(ctx))
		}

		// emit slash event
		types.EmitSlashEvent(ctx, pool.Id, pool.BundleProposal.Uploader, slashAmount)

		// Update the current lowest staker.
		k.updateLowestStaker(ctx, &pool)

		// Emit an invalid bundle event.
		types.EmitBundleInvalidEvent(ctx, &pool)

		// Update and return.
		pool.BundleProposal = &types.BundleProposal{
			NextUploader: pool.BundleProposal.NextUploader,
			FromHeight:   pool.BundleProposal.FromHeight,
			ToHeight:     pool.BundleProposal.FromHeight,
			CreatedAt:    uint64(ctx.BlockTime().Unix()),
		}

		k.SetPool(ctx, pool)

		return &types.MsgSubmitBundleProposalResponse{}, nil
	} else {
		// Throw error, should register bundle NO_QUORUM_BUNDLE instead.
		return nil, types.ErrQuorumNotReached
	}
}
