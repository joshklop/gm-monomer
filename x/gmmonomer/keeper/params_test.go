package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/joshklop/gm-monomer/testutil/keeper"
	"github.com/joshklop/gm-monomer/x/gmmonomer/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.GmmonomerKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
