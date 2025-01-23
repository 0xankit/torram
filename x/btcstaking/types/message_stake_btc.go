package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgStakeBtc{}

func NewMsgStakeBtc(creator string, txId string, vout uint32) *MsgStakeBtc {
	return &MsgStakeBtc{
		From: creator,
		TxId: txId,
		Vout: vout,
	}
}

func (msg *MsgStakeBtc) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
