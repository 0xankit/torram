package types

import fmt "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "btcstaking"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_btcstaking"

	// btc Denom
	BtcDenom = "trm"
)

var (
	ParamsKey = []byte("p_btcstaking")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// UTXOKey generates a unique key for storing UTXOs in the KVStore.
// It combines the transaction ID and output index (vout) into a single key.
func UTXOKey(txID string, vout uint32) []byte {
	return []byte(fmt.Sprintf("%s:%d", txID, vout))
}
