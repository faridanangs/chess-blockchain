package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/faridanangs/checkers/testutil/keeper"
	"github.com/faridanangs/checkers/x/checkers"
	"github.com/faridanangs/checkers/x/checkers/keeper"
	"github.com/faridanangs/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func setupCreateTodoServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())

	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateTodo(t *testing.T) {
	msgServer, k, context := setupCreateTodoServer(t)
	ctx := sdk.UnwrapSDKContext(context)

	todo := types.MsgCreateTodo{
		Creator: alice,
		Title:   "hello wrold",
		Text:    "lorem ipsum dolor sit amet",
	}

	resp, err := msgServer.CreateTodo(context, &todo)

	require.Nil(t, err)
	require.Equal(t, &types.MsgCreateTodoResponse{
		TodoIndex: "1",
	}, resp)

	todoResp, found := k.GetTodo(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.Todo{
		Index:   "1",
		Creator: alice,
		Title:   "hello wrold",
		Text:    "lorem ipsum dolor sit amet",
	}, todoResp)

	event := sdk.StringifyEvent(ctx.EventManager().ABCIEvents()[0])
	require.EqualValues(t, sdk.StringEvent{
		Type: "add-todo-event",
		Attributes: []sdk.Attribute{
			{Key: "index", Value: "1"},
			{Key: "title", Value: "hello wrold"},
			{Key: "text", Value: "lorem ipsum dolor sit amet"},
			{Key: "creator", Value: alice},
		},
	}, event)

}
