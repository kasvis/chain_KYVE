package v0_2_0

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	registrykeeper "github.com/KYVENetwork/chain/x/registry/keeper"
)

func CreateUpgradeHandler(
	registryKeeper *registrykeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		// Fetch the unbonding state.
		unbondingState, found := registryKeeper.GetUnbondingState(ctx)
		if !found {
			return vm, nil
		}

		// Check if queue is currently empty.
		if unbondingState.LowIndex == unbondingState.HighIndex {
			return vm, nil
		}

		// Iterate over queue and handle all entries.
		for i := unbondingState.LowIndex; i < unbondingState.HighIndex; i++ {
			// Fetch the current entry.
			unbondingEntry, found := registryKeeper.GetUnbondingEntries(ctx, i)
			if !found {
				continue
			}

			// Ensure that we're only handling delegation unbondings.
			if unbondingEntry.Delegator != unbondingEntry.Staker {
				// Transfer tokens from registry to delegator.
				err := registryKeeper.TransferToAddress(ctx, unbondingEntry.Delegator, unbondingEntry.Amount)
				if err != nil {
					registryKeeper.PanicHalt(ctx, "Not enough money in registry module: "+err.Error())
				}
			}
		}

		// Return.
		return vm, nil
	}
}
