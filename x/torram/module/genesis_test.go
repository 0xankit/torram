package torram_test

import (
	"testing"

	keepertest "torram/testutil/keeper"
	"torram/testutil/nullify"
	torram "torram/x/torram/module"
	"torram/x/torram/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TorramKeeper(t)
	torram.InitGenesis(ctx, k, genesisState)
	got := torram.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
