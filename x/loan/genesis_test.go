package loan_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/username/cosmos_sdk_demo/testutil/keeper"
	"github.com/username/cosmos_sdk_demo/testutil/nullify"
	"github.com/username/cosmos_sdk_demo/x/loan"
	"github.com/username/cosmos_sdk_demo/x/loan/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoanKeeper(t)
	loan.InitGenesis(ctx, *k, genesisState)
	got := loan.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
