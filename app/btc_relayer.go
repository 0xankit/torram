package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	btccontroller "torram/btc_controller"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/cometbft/cometbft/proto/tendermint/types"
)

func (app *App) StartRelayer() {
	go func() {
		// Graceful shutdown handling
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-sigs
			cancel()
		}()

		// btc relayer
		btcRelayer := btccontroller.NewBtcRelayer(
			&rpcclient.ConnConfig{
				Host:         "localhost:18443",
				User:         "yourusername",
				Pass:         "yourpassword",
				DisableTLS:   true,
				Params:       "regtest",
				HTTPPostMode: true,
			},
		)

		log.Println("Starting Bitcoin relayer...")

		err := monitorBitcoin(ctx, app, btcRelayer)
		if err != nil {
			log.Printf("Relayer error: %v", err)
		}

		log.Println("Bitcoin relayer stopped.")
	}()
}

func monitorBitcoin(ctx context.Context, app *App, btcRelayer *btccontroller.BTCRelayer) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			// Poll Bitcoin node and process transactions
			err := pollBitcoinNode(app, btcRelayer)
			if err != nil {
				log.Printf("Error polling Bitcoin node: %v", err)
			}
			time.Sleep(10 * time.Second) // Poll every 10 seconds
		}
	}
}

// pollBitcoinNode polls the Bitcoin node for new transactions and processes them
// using the app's handler
func pollBitcoinNode(app *App, btcRelayer *btccontroller.BTCRelayer) error {
	// List unspent outputs from Torram node
	sdkCtx := app.BaseApp.NewUncachedContext(false, types.Header{})
	unspentOutputs := app.BtcstakingKeeper.GetAllUTXOs(sdkCtx)
	for _, utxo := range unspentOutputs {
		if utxo.TxId == "" {
			continue
		}
		// Get raw transaction from Bitcoin node
		tx, err := btcRelayer.GetRawTransaction(utxo.TxId)
		if err != nil {
			return err
		}

		// Parse OP_RETURN data
		data := btccontroller.ParseOpReturn(tx)
		if data != "" {
			// Process OP_RETURN data
			err := btccontroller.ProcessOpReturnData(data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
