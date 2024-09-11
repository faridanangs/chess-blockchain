package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/faridanangs/checkers/x/checkers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TodoAll(goCtx context.Context, req *types.QueryAllTodoRequest) (*types.QueryAllTodoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var todos []types.Todo
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	todoStore := prefix.NewStore(store, types.KeyPrefix(types.TodoKeyPrefix))

	pageRes, err := query.Paginate(todoStore, req.Pagination, func(key []byte, value []byte) error {
		var todo types.Todo
		if err := k.cdc.Unmarshal(value, &todo); err != nil {
			return err
		}

		todos = append(todos, todo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTodoResponse{Todo: todos, Pagination: pageRes}, nil
}

func (k Keeper) Todo(goCtx context.Context, req *types.QueryGetTodoRequest) (*types.QueryGetTodoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTodo(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTodoResponse{Todo: val}, nil
}
