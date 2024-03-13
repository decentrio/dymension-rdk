package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMinter returns a new Minter object with the given epoch
// provisions values.
func NewMinter(inflationRate sdk.Dec) Minter {
	return Minter{
		CurrentInflationRate: inflationRate,
	}
}

// InitialMinter returns an initial Minter object.
func InitialMinter() Minter {
	return NewMinter(sdk.NewDec(0))
}

// DefaultInitialMinter returns a default initial Minter object for a new chain.
func DefaultInitialMinter() Minter {
	return InitialMinter()
}

// validate minter.
func ValidateMinter(minter Minter) error {
	return validateInflationRate(minter.CurrentInflationRate)
}
