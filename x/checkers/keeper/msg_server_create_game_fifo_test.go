package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faridanangs/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestCreate3GamesHasSavedInfo(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	// Second game
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
	})

	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        3,
		FifoHeadIndex: "1",
		FifoTailIndex: "2",
	}, systemInfo)

	game1, _ := keeper.GetStoredGame(ctx, "1")
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       bob,
		Red:         carol,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   uint64(0),
		BeforeIndex: "-1",
		AfterIndex:  "2",
	}, game1)

	game2, found := keeper.GetStoredGame(ctx, "2")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       carol,
		Red:         alice,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   uint64(0),
		BeforeIndex: "1",
		AfterIndex:  "-1",
	}, game2)

	// Third game
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
	})
	systemInfo3, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        4,
		FifoHeadIndex: "1",
		FifoTailIndex: "3",
	}, systemInfo3)

	game3, found := keeper.GetStoredGame(ctx, "3")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "3",
		Board:       "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "b",
		Black:       alice,
		Red:         bob,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   uint64(0),
		BeforeIndex: "2",
		AfterIndex:  "-1",
	}, game3)

	fmt.Println(game3, "game1")
	fmt.Println(systemInfo3, "system")

}

func TestPlayMove2Games2MovesHasSavedFifo(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForPlayMove(t)
	ctx := sdk.UnwrapSDKContext(context)
	msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
	})

	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})

	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   carol,
		GameIndex: "2",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})

	systemInfo1, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:        3,
		FifoHeadIndex: "1",
		FifoTailIndex: "2",
	}, systemInfo1)

	game1, found := keeper.GetStoredGame(ctx, "1")

	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "1",
		Board:       "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "r",
		Black:       bob,
		Red:         carol,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   uint64(1),
		BeforeIndex: "-1",
		AfterIndex:  "2",
	}, game1)

	game2, found := keeper.GetStoredGame(ctx, "2")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:       "2",
		Board:       "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:        "r",
		Black:       carol,
		Red:         alice,
		Winner:      "*",
		Deadline:    types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
		MoveCount:   uint64(1),
		BeforeIndex: "1",
		AfterIndex:  "-1",
	}, game2)
}
