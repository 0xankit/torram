package btcstaking

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "torram/api/torram/btcstaking"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "GetStakedBtc",
					Use:            "get-staked-btc [trx-id] [vout]",
					Short:          "Query getStakedBTC",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "trxId"}, {ProtoField: "vout"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "StakeBtc",
					Use:            "stake-btc [tx-id] [vout]",
					Short:          "Send a stakeBTC tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tx_id"}, {ProtoField: "vout"}},
				},
				{
					RpcMethod:      "UnstakeBtc",
					Use:            "unstake-btc [tx-id] [vout]",
					Short:          "Send a unstakeBTC tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tx_id"}, {ProtoField: "vout"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
