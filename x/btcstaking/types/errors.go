package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/btcstaking module sentinel errors
var (
	ErrUTXONotFound      = sdkerrors.Register(ModuleName, 1100, "UTXO not found")
	ErrUTXOAlreadyStaked = sdkerrors.Register(ModuleName, 1101, "UTXO is already staked")
	ErrUTXONotStaked     = sdkerrors.Register(ModuleName, 1102, "UTXO is not staked")
	ErrInvalidSigner     = sdkerrors.Register(ModuleName, 1103, "invalid signer")
)

// x/btcstaking emit Event
var (
	// btc_staked event
	EventTypeStake   = "btc_staked"
	AttributeKeyTxID = "tx_id"
	AttributeUtxo    = "utxo"

	// btc_unstaked event
	EventTypeUnstake = "btc_unstaked"
)
