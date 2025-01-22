package keeper

import (
	"torram/x/btcstaking/types"
)

var _ types.QueryServer = Keeper{}
