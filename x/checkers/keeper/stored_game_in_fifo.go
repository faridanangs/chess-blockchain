package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faridanangs/checkers/x/checkers/types"
)

func (k Keeper) RemoveFromFifo(ctx sdk.Context, game *types.StoredGame, info *types.SystemInfo) {
	if game.BeforeIndex != types.NoFifoIndex {
		beforeElemen, found := k.GetStoredGame(ctx, game.BeforeIndex)
		if !found {
			panic("Element before in fifo was not found")
		}

		beforeElemen.AfterIndex = game.AfterIndex
		k.SetStoredGame(ctx, beforeElemen)

		if game.AfterIndex == types.NoFifoIndex {
			info.FifoTailIndex = beforeElemen.Index
		}
	} else if info.FifoHeadIndex == game.Index {
		info.FifoHeadIndex = game.AfterIndex
	}

	if game.AfterIndex != types.NoFifoIndex {
		afterElement, found := k.GetStoredGame(ctx, game.AfterIndex)
		if !found {
			panic("Element after in fifo was not found")
		}

		afterElement.BeforeIndex = game.BeforeIndex
		k.SetStoredGame(ctx, afterElement)

		if game.BeforeIndex == types.NoFifoIndex {
			info.FifoHeadIndex = afterElement.Index
		}
	} else if info.FifoTailIndex == game.Index {
		info.FifoTailIndex = game.BeforeIndex
	}

	game.BeforeIndex = types.NoFifoIndex
	game.AfterIndex = types.NoFifoIndex
}

func (k Keeper) SendToFifoTail(ctx sdk.Context, game *types.StoredGame, info *types.SystemInfo) {
	if info.FifoHeadIndex == types.NoFifoIndex && info.FifoTailIndex == types.NoFifoIndex {
		game.BeforeIndex = types.NoFifoIndex
		game.AfterIndex = types.NoFifoIndex
		info.FifoHeadIndex = game.Index
		info.FifoTailIndex = game.Index

	} else if info.FifoHeadIndex == types.NoFifoIndex || info.FifoTailIndex == types.NoFifoIndex {
		panic("Fifo should have both head and tail or none")
	} else if info.FifoTailIndex == game.Index {
		// kosokng
	} else {
		k.RemoveFromFifo(ctx, game, info)

		currentTail, found := k.GetStoredGame(ctx, info.FifoTailIndex)

		if !found {
			panic("Current fifo tail was not found")
		}

		currentTail.AfterIndex = game.Index
		k.SetStoredGame(ctx, currentTail)

		game.BeforeIndex = currentTail.Index
		info.FifoTailIndex = game.Index
	}
}
