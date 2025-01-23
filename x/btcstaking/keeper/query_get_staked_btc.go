package keeper

import (
	"context"

	"torram/x/btcstaking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetStakedBtc(goCtx context.Context, req *types.QueryGetStakedBtcRequest) (*types.QueryGetStakedBtcResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Fetch the UTXO from the store
	utxo, found := k.GetUTXO(ctx, req.TrxId, req.Vout)
	if !found {
		return nil, types.ErrUTXONotFound
	}

	return &types.QueryGetStakedBtcResponse{Utxo: utxo}, nil
}
