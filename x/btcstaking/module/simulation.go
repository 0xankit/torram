package btcstaking

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"torram/testutil/sample"
	btcstakingsimulation "torram/x/btcstaking/simulation"
	"torram/x/btcstaking/types"
)

// avoid unused import issue
var (
	_ = btcstakingsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgStakeBtc = "op_weight_msg_stake_btc"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStakeBtc int = 100

	opWeightMsgUnstakeBtc = "op_weight_msg_unstake_btc"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnstakeBtc int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	btcstakingGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&btcstakingGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgStakeBtc int
	simState.AppParams.GetOrGenerate(opWeightMsgStakeBtc, &weightMsgStakeBtc, nil,
		func(_ *rand.Rand) {
			weightMsgStakeBtc = defaultWeightMsgStakeBtc
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStakeBtc,
		btcstakingsimulation.SimulateMsgStakeBtc(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnstakeBtc int
	simState.AppParams.GetOrGenerate(opWeightMsgUnstakeBtc, &weightMsgUnstakeBtc, nil,
		func(_ *rand.Rand) {
			weightMsgUnstakeBtc = defaultWeightMsgUnstakeBtc
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnstakeBtc,
		btcstakingsimulation.SimulateMsgUnstakeBtc(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgStakeBtc,
			defaultWeightMsgStakeBtc,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				btcstakingsimulation.SimulateMsgStakeBtc(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUnstakeBtc,
			defaultWeightMsgUnstakeBtc,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				btcstakingsimulation.SimulateMsgUnstakeBtc(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
