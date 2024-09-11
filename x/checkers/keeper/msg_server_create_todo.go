package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faridanangs/checkers/x/checkers/types"
)

func (k msgServer) CreateTodo(goCtx context.Context, msg *types.MsgCreateTodo) (*types.MsgCreateTodoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	createTodo := types.Todo{
		Index:   "1",
		Creator: msg.Creator,
		Title:   msg.Title,
		Text:    msg.Text,
	}

	k.Keeper.SetTodo(ctx, createTodo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.AddTodoEvent,
			sdk.Attribute{Key: "index", Value: "1"},
			sdk.Attribute{Key: "title", Value: msg.Title},
			sdk.Attribute{Key: "text", Value: msg.Text},
			sdk.Attribute{Key: "creator", Value: msg.Creator},
		))

	return &types.MsgCreateTodoResponse{
		TodoIndex: "1",
	}, nil
}
