package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faridanangs/checkers/x/checkers/testutil"
	"github.com/faridanangs/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestPlayMoveUpToWinner(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameForPlayMove(t)
	ctx := sdk.UnwrapSDKContext(context)

	testutil.PlayAllMoves(t, msgServer, context, "1", bob, carol, testutil.Game1Moves)

	systemInfo, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{NextId: 2}, systemInfo)

	game, found := keeper.GetStoredGame(ctx, "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index:  "1",
		Board:  "",
		Turn:   "b",
		Black:  bob,
		Red:    carol,
		Winner: "b",
	}, game)

	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())

	require.Len(t, events, 41)
	event := events[40]
	require.Equal(t, event.Type, "move-played")

	require.EqualValues(t, []sdk.Attribute{
		{Key: "creator", Value: bob},
		{Key: "game-index", Value: "1"},
		{Key: "captured-x", Value: "2"},
		{Key: "captured-y", Value: "5"},
		{Key: "winner", Value: "b"},
		{Key: "board", Value: "*b*b****|**b*b***|*****b**|********|***B****|********|*****b**|********"},
	}, event.Attributes)

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

	fmt.Println(sdk.StringifyEvents(ctx.EventManager().ABCIEvents()), "eventManager")
}
