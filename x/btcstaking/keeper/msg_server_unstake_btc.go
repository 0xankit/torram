package keeper

import (
	"context"

	"torram/x/btcstaking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UnstakeBtc(goCtx context.Context, msg *types.MsgUnstakeBtc) (*types.MsgUnstakeBtcResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the UTXO exists
	utxo, found := k.Keeper.GetUTXO(ctx, msg.TxId, msg.Vout)
	if !found {
		return nil, types.ErrUTXONotFound
	}

	// Check if the UTXO is not staked
	if !utxo.IsStaked {
		return nil, types.ErrUTXONotStaked
	}

	// Unstake the UTXO
	utxo.IsStaked = false
	k.Keeper.SetUTXO(ctx, utxo)

	// TODO: can use hooks to unstake after it is saved in the store instead of doing it here
	// burn the UTXO with the staking amount
	var burnCoins sdk.Coins
	burnCoins = burnCoins.Add(utxo.Amount)
	if err := k.Keeper.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins); err != nil {
		return nil, err
	}

	// emit Event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUnstake,
			sdk.NewAttribute(types.AttributeKeyTxID, msg.TxId),
			sdk.NewAttribute(types.AttributeUtxo, utxo.String()),
		),
	)

	return &types.MsgUnstakeBtcResponse{}, nil
}
