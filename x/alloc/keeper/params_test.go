package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/nghuyenthevinh2000/nebula/testutil/keeper"
	"github.com/nghuyenthevinh2000/nebula/x/alloc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AllocKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
