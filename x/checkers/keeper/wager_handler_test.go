package keeper_test

import (
	"context"
	"errors"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/faridanangs/checkers/testutil/keeper"
	"github.com/faridanangs/checkers/x/checkers"
	"github.com/faridanangs/checkers/x/checkers/keeper"
	"github.com/faridanangs/checkers/x/checkers/testutil"
	"github.com/faridanangs/checkers/x/checkers/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupKeeperForWagerHandler(t testing.TB) (keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankEscrowKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankEscrowKeeper(ctrl)
	k, ctx := keepertest.CheckersKeeperWithMocks(t, bankMock)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	context := sdk.WrapSDKContext(ctx)
	return *k, context, ctrl, bankMock
}

func TestWagerHandlerCollectWrongNoBlack(t *testing.T) {
	keeper, context, ctrl, _ := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "black address is invalid: : empty address string is not allowed", r)
	}()
	keeper.CollectWager(ctx, &types.StoredGame{
		MoveCount: 0,
	})
}

func TestWagerHandlerCollectFailedNoMove(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	black, _ := sdk.AccAddressFromBech32(alice)
	escrow.EXPECT().
		SendCoinsFromAccountToModule(ctx, black, types.ModuleName, gomock.Any()).
		Return(errors.New("oops"))
	err := keeper.CollectWager(ctx, &types.StoredGame{
		Black:     alice,
		MoveCount: 0,
		Wager:     45,
	})
	require.NotNil(t, err)
	require.EqualError(t, err, "black cannot pay the wager: oops")
}

func TestWagerHandlerCollectNoMove(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectPay(context, alice, 45)
	err := keeper.CollectWager(ctx, &types.StoredGame{
		Black:     alice,
		MoveCount: 0,
		Wager:     45,
	})
	require.Nil(t, err)
}

func TestWagerHandlerPayWrongEscrowFailed(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	black, _ := sdk.AccAddressFromBech32(alice)
	escrow.EXPECT().
		SendCoinsFromModuleToAccount(ctx, types.ModuleName, black, gomock.Any()).
		Times(1).
		Return(errors.New("oops"))
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, r, "cannot pay winnings to winner: oops")
	}()
	keeper.MustPayWinnings(ctx, &types.StoredGame{
		Black:     alice,
		Red:       bob,
		Winner:    "b",
		MoveCount: 1,
		Wager:     45,
	})
}

func TestWagerHandlerPayEscrowCalledTwoMoves(t *testing.T) {
	keeper, context, ctrl, escrow := setupKeeperForWagerHandler(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	escrow.ExpectRefund(context, alice, 90)
	keeper.MustPayWinnings(ctx, &types.StoredGame{
		Black:     alice,
		Red:       bob,
		Winner:    "b",
		MoveCount: 2,
		Wager:     45,
	})
}

func TestPlayMoveCalledBank(t *testing.T) {
	msgServer, _, context, ctrl, escrow := setupMsgServerWithOneGameForPlayMoveWithMock(t)
	defer ctrl.Finish()
	escrow.ExpectPay(context, bob, 45).Times(1)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
}

func TestPlayMoveUpToWinnerCalledBank(t *testing.T) {
	msgServer, _, context, ctrl, escrow := setupMsgServerWithOneGameForPlayMoveWithMock(t)
	defer ctrl.Finish()
	payBob := escrow.ExpectPay(context, bob, 45).Times(1)
	payCarol := escrow.ExpectPay(context, carol, 45).Times(1).After(payBob)
	escrow.ExpectRefund(context, bob, 90).Times(1).After(payCarol)

	testutil.PlayAllMoves(t, msgServer, context, "1", bob, carol, testutil.Game1Moves)
}

func TestForfeitPlayedOnceCalledBank(t *testing.T) {
	msgServer, keeper, context, ctrl, escrow := setupMsgServerWithOneGameForPlayMoveWithMock(t)
	ctx := sdk.UnwrapSDKContext(context)
	defer ctrl.Finish()
	pay := escrow.ExpectPay(context, bob, 45).Times(1)
	escrow.ExpectRefund(context, bob, 45).Times(1).After(pay)
	msgServer.PlayMove(context, &types.MsgPlayMove{
		Creator:   bob,
		GameIndex: "1",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	game1, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	game1.Deadline = types.FormatDeadline(ctx.BlockTime().Add(time.Duration(-1)))
	keeper.SetStoredGame(ctx, game1)
	keeper.ForfeitExpiredGames(context)
}
