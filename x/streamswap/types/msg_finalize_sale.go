package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFinalizeSale{}

func (msg *MsgFinalizeSale) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.ErrInvalidRequest.Wrapf("Invalid sender address (%s)", err)
	}
	return nil
}

func (msg *MsgFinalizeSale) Validate(now, end time.Time) error {
	if now.Before(end) {
		return errors.ErrInvalidRequest.Wrapf("You can exit the Sale only once the Sale ended (%s)", end)
	}
	return nil
}

func (msg *MsgFinalizeSale) GetSigners() []sdk.AccAddress {
	a, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{a}
}
