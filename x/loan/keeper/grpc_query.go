package keeper

import (
	"github.com/username/cosmos_sdk_demo/x/loan/types"
)

var _ types.QueryServer = Keeper{}
