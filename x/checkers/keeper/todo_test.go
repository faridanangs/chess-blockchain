package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/faridanangs/checkers/testutil/keeper"
	"github.com/faridanangs/checkers/testutil/nullify"
	"github.com/faridanangs/checkers/x/checkers/keeper"
	"github.com/faridanangs/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTodo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Todo {
	items := make([]types.Todo, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTodo(ctx, items[i])
	}
	return items
}

func TestTodoGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNTodo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTodo(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTodoRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNTodo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTodo(ctx,
			item.Index,
		)
		_, found := keeper.GetTodo(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTodoGetAll(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNTodo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTodo(ctx)),
	)
}
