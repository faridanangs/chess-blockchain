package keeper_test

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/faridanangs/checkers/testutil/keeper"
	"github.com/faridanangs/checkers/x/checkers"
	"github.com/faridanangs/checkers/x/checkers/keeper"
	"github.com/faridanangs/checkers/x/checkers/testutil"
	"github.com/faridanangs/checkers/x/checkers/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	server, k, context, _, escrow := setupMsgServerCreateGameWithMock(t)
	escrow.ExpectAny(context)
	return server, k, context
}

func setupMsgServerCreateGameWithMock(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankEscrowKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankEscrowKeeper(ctrl)
	k, ctx := keepertest.CheckersKeeperWithMocks(t, bankMock)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	return server, *k, context, ctrl, bankMock
}

func TestCreateGame(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	ctx := sdk.UnwrapSDKContext(context)

	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Wager:   45,
	})

	require.Nil(t, err)
	require.EqualValues(t, &types.MsgCreateGameResponse{
		GameIndex: "1",
	}, createResponse)

	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "new-game-created",
		Attributes: []sdk.Attribute{
			{
				Key: "creator", Value: alice,
			},
			{
				Key: "game-index", Value: "1",
			},
			{
				Key: "black", Value: bob,
			},
			{
				Key: "red", Value: carol,
			},
			{
				Key: "wager", Value: "45",
			},
		},
	}, event)

	// Block header details
	fmt.Println(ctx.BlockHeader().AppHash, "Application Hash from Block Header")
	fmt.Println(ctx.BlockHeader().ChainID, "Chain ID from Block Header")
	fmt.Println(ctx.BlockHeader().ConsensusHash, "Consensus Hash from Block Header")
	fmt.Println(ctx.BlockHeader().DataHash, "Data Hash from Block Header")
	fmt.Println(ctx.BlockHeader().EvidenceHash, "Evidence Hash from Block Header")
	fmt.Println(ctx.BlockHeader().Height, "Block Height")
	fmt.Println(ctx.BlockHeader().LastBlockId, "Last Block ID")
	fmt.Println(ctx.BlockHeader().LastCommitHash, "Last Commit Hash from Block Header")
	fmt.Println(ctx.BlockHeader().LastResultsHash, "Last Results Hash from Block Header")
	fmt.Println(ctx.BlockHeader().NextValidatorsHash, "Next Validators Hash from Block Header")
	fmt.Println(ctx.BlockHeader().ProposerAddress, "Proposer Address from Block Header")
	fmt.Println(ctx.BlockHeader().Time, "Block Time")
	fmt.Println(ctx.BlockHeader().ValidatorsHash, "Validators Hash from Block Header")
	fmt.Println(ctx.BlockHeader().Version, "Block Version")
	fmt.Println(ctx.BlockHeader().LastBlockId.Hash, "Hash of Last Block ID")
	fmt.Println(ctx.BlockHeader().Version.App, "Application Version from Block Header")
	fmt.Println(ctx.BlockHeader().Version.Block, "Block Version from Block Header")

	// Context-related details
	fmt.Println(ctx.BlockHeight(), "Block Height")
	fmt.Println(ctx.BlockTime(), "Block Time")
	fmt.Println(ctx.ConsensusParams(), "Consensus Parameters")
	fmt.Println(ctx.ChainID(), "Chain ID")
	fmt.Println(ctx.Done(), "Done Channel in Context")
	fmt.Println(ctx.Err(), "Error in Context")

	// Gas-related details
	fmt.Println(ctx.GasMeter().GasConsumed(), "Gas Consumed by Transaction")
	fmt.Println(ctx.GasMeter().GasRemaining(), "Remaining Gas in Transaction")
	fmt.Println(ctx.GasMeter().String(), "Gas Meter as String")

	// KV Store related details
	fmt.Println(ctx.KVGasConfig().ReadCostFlat, "Flat Cost for Read Operation in KV Store")
	fmt.Println(ctx.KVGasConfig().ReadCostPerByte, "Cost per Byte for Read Operation in KV Store")
	fmt.Println(ctx.KVGasConfig().DeleteCost, "Cost for Delete Operation in KV Store")
	fmt.Println(ctx.KVGasConfig().HasCost, "Cost for Has Operation in KV Store")
	fmt.Println(ctx.KVGasConfig().IterNextCostFlat, "Flat Cost for Iteration in KV Store")
	fmt.Println(ctx.KVGasConfig().WriteCostFlat, "Flat Cost for Write Operation in KV Store")
	fmt.Println(ctx.KVGasConfig().WriteCostPerByte, "Cost per Byte for Write Operation in KV Store")
	fmt.Println(ctx.KVGasConfig(), "KV Store Gas Configuration")

	// MultiStore-related details
	fmt.Println(ctx.MultiStore().CacheMultiStore(), "Cache MultiStore")
	fmt.Println(ctx.MultiStore().CacheWrap(), "Cache Wrap of MultiStore")
	fmt.Println(ctx.MultiStore().GetStoreType(), "Store Type of MultiStore")
	fmt.Println(ctx.MultiStore().LatestVersion(), "Latest Version of MultiStore")
	fmt.Println(ctx.MultiStore().TracingEnabled(), "Is Tracing Enabled in MultiStore")

	// Additional details
	fmt.Println(ctx.HeaderHash().String(), "Header Hash as String")
	fmt.Println(ctx.IsCheckTx(), "Is Check Transaction")
	fmt.Println(ctx.IsZero(), "Is Context Zero")
	fmt.Println(ctx.Logger(), "Logger")
	fmt.Println(ctx.VoteInfos(), "Vote Infos")
	fmt.Println(ctx.TxBytes(), "Transaction Bytes")
	fmt.Println(ctx.IsReCheckTx(), "Is ReCheck Transaction")
	fmt.Println(ctx.MinGasPrices().String(), "Minimum Gas Prices as String")
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

func TestCreate1GameEmitted(t *testing.T) {
	msgSrvr, _, context := setupMsgServerCreateGame(t)
	msgSrvr.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)

	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)

	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "new-game-created",
		Attributes: []sdk.Attribute{
			{
				Key: "creator", Value: alice,
			},
			{
				Key: "game-index", Value: "1",
			},
			{
				Key: "black", Value: bob,
			},
			{
				Key: "red", Value: carol,
			},
		},
	}, event)

}
