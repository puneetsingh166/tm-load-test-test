package v040

import (
	v039crisis "github.com/puneetsingh166/tm-load-test/x/crisis/migrations/v039"
	v040crisis "github.com/puneetsingh166/tm-load-test/x/crisis/types"
)

// Migrate accepts exported v0.39 x/crisis genesis state and
// migrates it to v0.40 x/crisis genesis state. The migration includes:
//
// - Re-encode in v0.40 GenesisState.
func Migrate(crisisGenState v039crisis.GenesisState) *v040crisis.GenesisState {
	return &v040crisis.GenesisState{
		ConstantFee: crisisGenState.ConstantFee,
	}
}
