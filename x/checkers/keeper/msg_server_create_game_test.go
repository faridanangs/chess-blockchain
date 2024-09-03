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

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)

	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())

	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestCreateGame(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)

	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	require.Nil(t, err)
	require.EqualValues(t, &types.MsgCreateGameResponse{
		GameIndex: "1",
	}, createResponse)
}

func TestCreate1GameHasSaved(t *testing.T) {
	msgSrvr, keeper, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)

	msgSrvr.CreateGame(ctx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.Equal(t, types.SystemInfo{
		NextId: 2,
	}, systemInfo)

	storedGame, found1 := keeper.GetStoredGame(ctx, "1")
	require.True(t, found1)

	require.Equal(t, types.StoredGame{
		Index: "1",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: bob,
		Red:   carol,
	}, storedGame)
}

func TestCreate1GameGetAll(t *testing.T) {
	msgSrvr, keeper, contex := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(contex)

	msgSrvr.CreateGame(ctx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	games := keeper.GetAllStoredGame(ctx)

	require.Equal(t, types.StoredGame{
		Index: "1",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: bob,
		Red:   carol,
	}, games[0])
}

func TestCreateGameRedAddressBad(t *testing.T) {
	msgSrvr, _, contex := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(contex)

	resp, err := msgSrvr.CreateGame(ctx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     "notanaddress",
	})

	require.Nil(t, resp)
	require.Equal(t,
		"red address is invalid: notanaddress: decoding bech32 failed: invalid separator index -1",
		err.Error())

}

func TestCreate3Games(t *testing.T) {
	msgSrvr, _, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	createResponse2, err2 := msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: bob,
		Black:   carol,
		Red:     alice,
	})
	require.Nil(t, err2)

	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "2",
	}, *createResponse2)

	createResponse3, err3 := msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: carol,
		Black:   alice,
		Red:     bob,
	})
	require.Nil(t, err3)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "3",
	}, *createResponse3)
}
