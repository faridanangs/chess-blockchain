package checkers_test

import (
	"testing"

	keepertest "github.com/faridanangs/checkers/testutil/keeper"
	"github.com/faridanangs/checkers/testutil/nullify"
	"github.com/faridanangs/checkers/x/checkers"
	"github.com/faridanangs/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextId: 52,
		},
		StoredGameList: []types.StoredGame{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		TodoList: []types.Todo{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, genesisState)
	got := checkers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredGameList, got.StoredGameList)
	require.ElementsMatch(t, genesisState.TodoList, got.TodoList)
	// this line is used by starport scaffolding # genesis/test/assert
}

func TestDefaultGenesisState_ExpectedInitialNextId(t *testing.T) {
	require.EqualValues(t,
		&types.GenesisState{
			StoredGameList: []types.StoredGame{},
			SystemInfo: types.SystemInfo{
				NextId:        uint64(1),
				FifoHeadIndex: "-1",
				FifoTailIndex: "-1",
			},
		},
		types.DefaultGenesis())
}
