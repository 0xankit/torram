package keeper

import (
	"torram/x/btcstaking/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// getKVStore returns a prefixed KVStore
func (k Keeper) getKVStore(ctx sdk.Context) prefix.Store {
	// Helper function to create a prefixed KVStore
	kvStore := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(kvStore, types.KeyPrefix(types.StoreKey))
}

// SetUTXO saves a UTXO in the store
func (k Keeper) SetUTXO(ctx sdk.Context, utxo types.UTXO) {
	store := k.getKVStore(ctx)
	b := k.cdc.MustMarshal(&utxo)
	store.Set(types.UTXOKey(utxo.TxId, utxo.Vout), b)
}

// GetUTXO retrieves a UTXO by txID and vout
func (k Keeper) GetUTXO(ctx sdk.Context, txID string, vout uint32) (types.UTXO, bool) {
	store := k.getKVStore(ctx)
	b := store.Get(types.UTXOKey(txID, vout))
	if b == nil {
		return types.UTXO{}, false
	}

	var utxo types.UTXO
	k.cdc.MustUnmarshal(b, &utxo)
	return utxo, true
}

// RemoveUTXO deletes a UTXO from the store
func (k Keeper) RemoveUTXO(ctx sdk.Context, txID string, vout uint32) {
	store := k.getKVStore(ctx)
	store.Delete(types.UTXOKey(txID, vout))
}

// GetAllUTXOs retrieves all UTXOs from the store
func (k Keeper) GetAllUTXOs(ctx sdk.Context) []types.UTXO {
	store := k.getKVStore(ctx)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var utxos []types.UTXO
	for ; iterator.Valid(); iterator.Next() {
		var utxo types.UTXO
		k.cdc.MustUnmarshal(iterator.Value(), &utxo)
		utxos = append(utxos, utxo)
	}

	return utxos
}
