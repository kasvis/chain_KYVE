package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/KYVENetwork/chain/x/registry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakersByPoolAndDelegator(goCtx context.Context, req *types.QueryStakersByPoolAndDelegatorRequest) (*types.QueryStakersByPoolAndDelegatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pool, found := k.GetPool(ctx, req.PoolId)
	if !found {
		return nil, sdkErrors.Wrapf(sdkErrors.ErrNotFound, types.ErrPoolNotFound.Error(), req.PoolId)
	}

	var stakers []types.DelegationForStakerResponse

	store := ctx.KVStore(k.storeKey)
	delegatorStore := prefix.NewStore(store, types.KeyPrefix(types.DelegatorKeyPrefix))

	pageRes, err := query.FilteredPaginate(delegatorStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var delegator types.Delegator

		if err := k.cdc.Unmarshal(value, &delegator); err != nil {
			return false, nil
		}

		if delegator.Delegator != req.Delegator {
			return false, nil
		}

		if delegator.Id != req.PoolId {
			return false, nil
		}

		if accumulate {

			f1 := F1Distribution{
				k:                k,
				ctx:              ctx,
				poolId:           pool.Id,
				stakerAddress:    delegator.Staker,
				delegatorAddress: delegator.Delegator,
			}

			delegationPoolData, _ := k.GetDelegationPoolData(ctx, pool.Id, delegator.Staker)

			stakers = append(stakers, types.DelegationForStakerResponse{
				Staker:                delegator.Staker,
				CurrentReward:         f1.getCurrentReward(),
				DelegationAmount:      delegator.DelegationAmount,
				TotalDelegationAmount: delegationPoolData.TotalDelegation,
				DelegatorCount:        delegationPoolData.DelegatorCount,
			})

		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryStakersByPoolAndDelegatorResponse{
		Delegator:  req.Delegator,
		Pool:       &pool,
		Stakers:    stakers,
		Pagination: pageRes,
	}, nil
}
