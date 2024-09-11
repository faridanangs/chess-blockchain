package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/faridanangs/checkers/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTodo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTodo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTodo{
				Creator: "invalid_address",
				Title:   "hello world",
				Text:    "hello world",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTodo{
				Creator: sample.AccAddress(),
				Title:   "hello world",
				Text:    "hello world",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
