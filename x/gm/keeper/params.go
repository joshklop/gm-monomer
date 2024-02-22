package keeper

import (
	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
	"github.com/joshklop/gm-monomer/x/gm/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams()
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
