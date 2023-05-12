package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBroadcastRandom{}

func NewMsgBroadcastRandom(index string, creator string, vrv string, multiplier uint64, proof string, pubk string, message string, parsedvrv uint64, floatvrv float64, finalvrv uint64, finalvrvfl float64) *MsgBroadcastRandom {
	return &MsgBroadcastRandom{
//		Creator:   creator,
//		Multiplier: multiplier,
		Index:     index,
		Creator:   creator,
		Vrv:       vrv,
		Multiplier:multiplier,
		Proof:     proof,
		Pubk:      pubk,
		Message:   message,
		Parsedvrv: parsedvrv,
		Floatvrv:  floatvrv,
		Finalvrv:  finalvrv,
		Finalvrvfl: finalvrvfl,
	}
}

func (msg *MsgBroadcastRandom) Route() string {
	return RouterKey
}

func (msg *MsgBroadcastRandom) Type() string {
	return "BroadcastRandom"
}

func (msg *MsgBroadcastRandom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBroadcastRandom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBroadcastRandom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
