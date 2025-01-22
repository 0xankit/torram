package btcstaking_test

import (
	"testing"

	keepertest "torram/testutil/keeper"
	"torram/testutil/nullify"
	btcstaking "torram/x/btcstaking/module"
	"torram/x/btcstaking/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BtcstakingKeeper(t)
	btcstaking.InitGenesis(ctx, k, genesisState)
	got := btcstaking.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
