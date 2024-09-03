package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faridanangs/checkers/x/checkers/rules"
	"github.com/faridanangs/checkers/x/checkers/testutil"
	"github.com/stretchr/testify/require"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
)

func GetStoredGame1() StoredGame {
	return StoredGame{
		Black: alice,
		Red:   bob,
		Index: "1",
		Board: rules.New().String(),
		Turn:  "b",
	}
}

func TestCanGetAddressBlack(t *testing.T) {
	alice, err := sdk.AccAddressFromBech32(alice)
	black, err2 := GetStoredGame1().GetBlackAddress()

	require.Equal(t, alice, black)
	require.Nil(t, err)
	require.Nil(t, err2)
}

func TestGetAddressWrongBlack(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4"
	black, err := storedGame.GetBlackAddress()

	require.Nil(t, black)
	require.EqualError(t,
		err,
		"black address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)")
	require.EqualError(t, storedGame.Validate(), err.Error())
	require.EqualError(t, storedGame.Validate(), err.Error())
}
func TestCanGetAddressRed(t *testing.T) {
	bob, err := sdk.AccAddressFromBech32(bob)
	red, err2 := GetStoredGame1().GetRedAddress()

	require.Equal(t, bob, red)
	require.Nil(t, err)
	require.Nil(t, err2)
}

func TestGetAddressWrongRed(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Red = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8j"
	red, err := storedGame.GetRedAddress()

	require.Nil(t, red)
	require.EqualError(t,
		err,
		"red address is invalid: cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8j: decoding bech32 failed: invalid checksum (expected xqhc8g got xqhc8j)")
	require.EqualError(t, storedGame.Validate(), err.Error())
	require.EqualError(t, storedGame.Validate(), err.Error())
}
