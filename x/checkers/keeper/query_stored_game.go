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

func (k Keeper) StoredGameAll(goCtx context.Context, req *types.QueryAllStoredGameRequest) (*types.QueryAllStoredGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedGames []types.StoredGame
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	storedGameStore := prefix.NewStore(store, types.KeyPrefix(types.StoredGameKeyPrefix))

	pageRes, err := query.Paginate(storedGameStore, req.Pagination, func(key []byte, value []byte) error {
		var storedGame types.StoredGame
		if err := k.cdc.Unmarshal(value, &storedGame); err != nil {
			return err
		}

		storedGames = append(storedGames, storedGame)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredGameResponse{StoredGame: storedGames, Pagination: pageRes}, nil
}

func (k Keeper) StoredGame(goCtx context.Context, req *types.QueryGetStoredGameRequest) (*types.QueryGetStoredGameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetStoredGame(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredGameResponse{StoredGame: val}, nil
}
