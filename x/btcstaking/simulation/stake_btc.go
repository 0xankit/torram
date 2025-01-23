package simulation

import (
	"math/rand"

	"torram/x/btcstaking/keeper"
	"torram/x/btcstaking/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgStakeBtc(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgStakeBtc{
			From: simAccount.Address.String(),
		}

		// TODO: Handling the StakeBtc simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "StakeBtc simulation not implemented"), nil, nil
	}
}
