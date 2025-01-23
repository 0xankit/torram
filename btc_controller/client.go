package btccontroller

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
)

type BTCRelayer struct {
	BitcoinClient *rpcclient.Client
}

func NewBtcRelayer(connConfig *rpcclient.ConnConfig) *BTCRelayer {
	client, err := rpcclient.New(connConfig, nil)
	if err != nil {
		panic(err)
	}
	return &BTCRelayer{
		BitcoinClient: client,
	}
}

// listUnspent returns all unspent outputs for an address
func (r *BTCRelayer) ListUnspent() ([]btcjson.ListUnspentResult, error) {
	unspentOutputs, err := r.BitcoinClient.ListUnspent()
	if err != nil {
		return nil, err
	}
	return unspentOutputs, nil
}

// getRawTransaction returns the raw transaction for a transaction hash
func (r *BTCRelayer) GetRawTransaction(txID string) (*wire.MsgTx, error) {
	hash, err := chainhash.NewHashFromStr(txID)
	if err != nil {
		return nil, err
	}
	tx, err := r.BitcoinClient.GetRawTransaction(hash)
	if err != nil {
		return nil, err
	}

	return tx.MsgTx(), nil
}

// parse OP_RETURN from a transaction
func ParseOpReturn(tx *wire.MsgTx) string {
	for _, txOut := range tx.TxOut {
		if len(txOut.PkScript) > 0 && txOut.PkScript[0] == 0x6a {
			return string(txOut.PkScript[2:])
		}
	}
	return ""
}

// Process OP_RETURN data
func ProcessOpReturnData(data string) error {
	// Process the data here
	return nil
}
