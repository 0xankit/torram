package keeper

import (
	"context"

	"torram/x/btcstaking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) StakeBtc(goCtx context.Context, msg *types.MsgStakeBtc) (*types.MsgStakeBtcResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the UTXO exists
	utxo, found := k.Keeper.GetUTXO(ctx, msg.TxId, msg.Vout)

	// Check if the UTXO is already staked
	if found && utxo.IsStaked {
		//TODO: Update stake logic
		return nil, types.ErrUTXOAlreadyStaked
	}
	// Stake the UTXO

	utxo = types.UTXO{
		TxId:     msg.TxId,
		Vout:     msg.Vout,
		Amount:   msg.GetAmount(),
		Address:  msg.From,
		IsStaked: true,
	}
	k.Keeper.SetUTXO(ctx, utxo)

	return &types.MsgStakeBtcResponse{}, nil
}
