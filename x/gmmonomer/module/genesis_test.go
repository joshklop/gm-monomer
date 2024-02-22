package gmmonomer_test

import (
	"testing"

	keepertest "github.com/joshklop/gm-monomer/testutil/keeper"
	"github.com/joshklop/gm-monomer/testutil/nullify"
	gmmonomer "github.com/joshklop/gm-monomer/x/gmmonomer/module"
	"github.com/joshklop/gm-monomer/x/gmmonomer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GmmonomerKeeper(t)
	gmmonomer.InitGenesis(ctx, k, genesisState)
	got := gmmonomer.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
