package participate_test

import (
	"testing"

	keepertest "github.com/nghuyenthevinh2000/nebula/testutil/keeper"
	"github.com/nghuyenthevinh2000/nebula/testutil/nullify"
	"github.com/nghuyenthevinh2000/nebula/x/participate"
	"github.com/nghuyenthevinh2000/nebula/x/participate/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ParticipateKeeper(t)
	participate.InitGenesis(ctx, *k, genesisState)
	got := participate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
