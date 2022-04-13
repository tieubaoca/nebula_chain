package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/nghuyenthevinh2000/nebula/x/alloc/types"
    "github.com/nghuyenthevinh2000/nebula/x/alloc/keeper"
    keepertest "github.com/nghuyenthevinh2000/nebula/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AllocKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
