package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/nghuyenthevinh2000/nebula/testutil/keeper"
	"github.com/nghuyenthevinh2000/nebula/x/participate/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ParticipateKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
