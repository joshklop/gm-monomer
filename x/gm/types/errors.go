package types

// DONTCOVER

import (
	sdkerrors "github.com/joshklop/monomer-cosmos-sdk/types/errors"
)

// x/gm module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)
