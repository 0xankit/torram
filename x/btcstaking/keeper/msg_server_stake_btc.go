package keeper

import (
	"context"

	"torram/x/btcstaking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakeBtc defines a method for staking a UTXO, If the UTXO is already staked, it returns an error. It emits a Stake event
func (k msgServer) StakeBtc(goCtx context.Context, msg *types.MsgStakeBtc) (*types.MsgStakeBtcResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the UTXO exists
	utxo, found := k.Keeper.GetUTXO(ctx, msg.TxId, msg.Vout)

	// Check if the UTXO is already staked
	if found && utxo.IsStaked {
		//TODO: Update stake logic
		return nil, types.ErrUTXOAlreadyStaked
	}
	// Record the UTXO

	utxo = types.UTXO{
		TxId:     msg.TxId,
		Vout:     msg.Vout,
		Amount:   msg.GetAmount(), // This is not original UTXO amount, but the amount to be staked
		Address:  msg.From,
		IsStaked: true,
	}
	k.Keeper.SetUTXO(ctx, utxo)

	// TODO: can use hooks to stake after it is saved in the store instead of doing it here
	// mint the UTXO with the staking amount
	var mintCoins sdk.Coins
	mintCoins = mintCoins.Add(msg.GetAmount())
	if err := k.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return nil, err
	}

	// emit Event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeStake,
			sdk.NewAttribute(types.AttributeKeyTxID, msg.TxId),
			sdk.NewAttribute(types.AttributeUtxo, utxo.String()),
		),
	)

	return &types.MsgStakeBtcResponse{}, nil
}
