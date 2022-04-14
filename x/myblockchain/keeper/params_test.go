package keeper_test

import (
	"testing"

	testkeeper "github.com/daniel-farina/myblockchain/testutil/keeper"
	"github.com/daniel-farina/myblockchain/x/myblockchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MyblockchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
