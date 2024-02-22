package keeper

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/joshklop/monomer-cosmos-sdk/codec"
	codectypes "github.com/joshklop/monomer-cosmos-sdk/codec/types"
	"github.com/joshklop/monomer-cosmos-sdk/runtime"
	sdk "github.com/joshklop/monomer-cosmos-sdk/types"
	authtypes "github.com/joshklop/monomer-cosmos-sdk/x/auth/types"
	govtypes "github.com/joshklop/monomer-cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	"github.com/joshklop/gm-monomer/x/gmmonomer/keeper"
	"github.com/joshklop/gm-monomer/x/gmmonomer/types"
)

func GmmonomerKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	k := keeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
