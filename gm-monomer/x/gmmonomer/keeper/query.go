package keeper

import (
	"github.com/joshklop/gm-monomer/x/gmmonomer/types"
)

var _ types.QueryServer = Keeper{}
