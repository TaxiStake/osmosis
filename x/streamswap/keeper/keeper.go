package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/osmosis-labs/osmosis/v11/x/streamswap"
)

type Keeper struct {
	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	bank       BankKeeper
	paramSpace paramtypes.Subspace
}

// NewKeeper constructs a new streamswap Keeper
func NewKeeper(storeKey storetypes.StoreKey, cdc codec.BinaryCodec, bank BankKeeper, paramSpace paramtypes.Subspace) Keeper {
	return Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		bank:       bank,
		paramSpace: paramSpace,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", streamswap.ModuleName))
}
