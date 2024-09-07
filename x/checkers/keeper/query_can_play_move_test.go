package keeper_test

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/faridanangs/checkers/testutil/keeper"
	"github.com/faridanangs/checkers/x/checkers"
	"github.com/faridanangs/checkers/x/checkers/keeper"
	"github.com/faridanangs/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func setUpCanPlayMove(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := testutil.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())

	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCanPlayMove(t *testing.T) {
	msgServer, k, context := setUpCanPlayMove(t)

	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	require.Nil(t, err)
	require.EqualValues(t, &types.MsgCreateGameResponse{
		GameIndex: "1",
	}, createResponse)

	resp, err := k.CanPlayMove(context, &types.QueryCanPlayMoveRequest{
		GameIndex: "1",
		Player:    "b",
		FromX:     0,
		FromY:     5,
		ToX:       1,
		ToY:       4,
	})

	fmt.Println(resp, "resp")
	fmt.Println(err, "error")
}
