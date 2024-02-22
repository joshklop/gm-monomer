package keeper_test

import (
	"testing"

	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
	testkeeper "github.com/joshklop/gm-monomer/testutil/keeper"
	"github.com/joshklop/gm-monomer/x/gm/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.GmKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
