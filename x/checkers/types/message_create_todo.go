package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateTodo = "create_todo"

var _ sdk.Msg = &MsgCreateTodo{}

func NewMsgCreateTodo(creator string, title string, text string) *MsgCreateTodo {
	return &MsgCreateTodo{
		Creator: creator,
		Title:   title,
		Text:    text,
	}
}

func (msg *MsgCreateTodo) Route() string {
	return RouterKey
}

func (msg *MsgCreateTodo) Type() string {
	return TypeMsgCreateTodo
}

func (msg *MsgCreateTodo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTodo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTodo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
