package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/cosmonaut/scavenge/testutil/keeper"
	"github.com/cosmonaut/scavenge/testutil/nullify"
	"github.com/cosmonaut/scavenge/x/scavenge/keeper"
	"github.com/cosmonaut/scavenge/x/scavenge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNScavenge(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Scavenge {
	items := make([]types.Scavenge, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetScavenge(ctx, items[i])
	}
	return items
}

func TestScavengeGet(t *testing.T) {
	keeper, ctx := keepertest.ScavengeKeeper(t)
	items := createNScavenge(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetScavenge(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestScavengeRemove(t *testing.T) {
	keeper, ctx := keepertest.ScavengeKeeper(t)
	items := createNScavenge(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveScavenge(ctx,
			item.Index,
		)
		_, found := keeper.GetScavenge(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestScavengeGetAll(t *testing.T) {
	keeper, ctx := keepertest.ScavengeKeeper(t)
	items := createNScavenge(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllScavenge(ctx)),
	)
}
