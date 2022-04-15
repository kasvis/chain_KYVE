package keeper

import (
	"context"

	"github.com/KYVENetwork/chain/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// StakePool handles the logic of an SDK message that allows protocol nodes to stake in a specified pool.
func (k msgServer) StakePool(goCtx context.Context, msg *types.MsgStakePool) (*types.MsgStakePoolResponse, error) {
	// Unwrap context and attempt to fetch the pool.
	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, found := k.GetPool(ctx, msg.Id)

	// Error if the pool isn't found.
	if !found {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), msg.Id)
	}

	// Check if the sender is already a staker.
	staker, stakerExists := k.GetStaker(ctx, msg.Creator, msg.Id)

	if stakerExists {
		staker.Amount += msg.Amount
		k.SetStaker(ctx, staker)
	} else {
		// Check if we have reached the maximum number of stakers.
		// If we are staking more than the lowest staker, remove them.
		if len(pool.Stakers) == types.MaxStakers {
			lowestStaker, _ := k.GetStaker(ctx, pool.LowestStaker, msg.Id)

			if msg.Amount > lowestStaker.Amount {
				// Transfer tokens from this module to the lowest staker.
				err := k.TransferToAddress(ctx, lowestStaker.Account, lowestStaker.Amount)
				if err != nil {
					return nil, err
				}

				// Emit an unstake event.
				types.EmitUnstakeEvent(ctx, msg.Id, lowestStaker.Account, lowestStaker.Amount)

				// Remove lowest staker.
				k.removeStaker(ctx, &pool, &lowestStaker)
			} else {
				return nil, sdkErrors.Wrapf(sdkErrors.ErrLogic, types.ErrStakeTooLow.Error(), lowestStaker.Amount)
			}
		}

		pool.Stakers = append(pool.Stakers, msg.Creator)
		if staker.Commission == "" {
			k.SetStaker(ctx, types.Staker{
				Account:    msg.Creator,
				PoolId:     msg.Id,
				Amount:     msg.Amount,
				Commission: types.DefaultCommission,
			})
		} else {
			k.SetStaker(ctx, types.Staker{
				Account: msg.Creator,
				PoolId:  msg.Id,
				Amount:  msg.Amount,
			})
		}
	}

	// Transfer tokens from sender to this module.
	err := k.transferToRegistry(ctx, msg.Creator, msg.Amount)
	if err != nil {
		return nil, err
	}

	// Event a stake event.
	types.EmitStakeEvent(ctx, msg.Creator, msg.Id, msg.Amount)

	// Update and return.
	pool.TotalStake += msg.Amount
	k.updateLowestStaker(ctx, &pool)
	k.SetPool(ctx, pool)

	return &types.MsgStakePoolResponse{}, nil
}
