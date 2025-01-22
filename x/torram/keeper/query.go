package keeper

import (
	"torram/x/torram/types"
)

var _ types.QueryServer = Keeper{}
