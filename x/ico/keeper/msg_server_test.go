package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/nghuyenthevinh2000/nebula/x/ico/types"
    "github.com/nghuyenthevinh2000/nebula/x/ico/keeper"
    keepertest "github.com/nghuyenthevinh2000/nebula/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IcoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
