package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/nghuyenthevinh2000/nebula/x/participate/types"
    "github.com/nghuyenthevinh2000/nebula/x/participate/keeper"
    keepertest "github.com/nghuyenthevinh2000/nebula/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ParticipateKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
