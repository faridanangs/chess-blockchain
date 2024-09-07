package keeper

import (
	"context"
	"fmt"

	"github.com/faridanangs/checkers/x/checkers/rules"
	"github.com/faridanangs/checkers/x/checkers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CanPlayMove(goCtx context.Context, req *types.QueryCanPlayMoveRequest) (*types.QueryCanPlayMoveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedGame types.StoredGame

	// TODO: Process the query
	isBlack := req.Player == "b"
	isRed := req.Player == "r"

	var player rules.Player

	if isBlack && isRed {
		player = rules.StringPieces[storedGame.Turn].Player
	} else if isBlack {
		player = rules.BLACK_PLAYER
	} else if isRed {
		player = rules.RED_PLAYER
	} else {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   fmt.Sprintf("%s: %s", "message creator is not a player", req.Player),
		}, nil
	}

	game, err := storedGame.ParseGame()
	if err != nil {
		return nil, err
	}

	if !game.TurnIs(player) {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   fmt.Sprintf("%s: %s", "player tried to play out of turn", player.Color),
		}, nil
	}

	_, moveErr := game.Move(
		rules.Pos{
			X: int(req.FromX),
			Y: int(req.FromY),
		},
		rules.Pos{
			X: int(req.ToX),
			Y: int(req.ToY),
		},
	)

	if moveErr != nil {
		return &types.QueryCanPlayMoveResponse{
			Possible: false,
			Reason:   fmt.Sprintf("%s: %s", "wrong move", moveErr.Error()),
		}, nil
	}

	return &types.QueryCanPlayMoveResponse{
		Possible: true,
		Reason:   "ok",
	}, nil
}
