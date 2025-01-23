package btcstaking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"torram/x/btcstaking/keeper"
	"torram/x/btcstaking/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// to inclue the initial state of staking pool
	for _, utxo := range genState.Utxos {
		k.SetUTXO(ctx, *utxo)
	}

	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// inclue utxos in genesis
	allUtxos := k.GetAllUTXOs(ctx)
	genesis.Utxos = make([]*types.UTXO, len(allUtxos))
	for i, utxo := range allUtxos {
		genesis.Utxos[i] = &utxo // Take the address of each UTXO
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
