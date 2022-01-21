package poa

import (
	"github.com/youngjun1714/sdk/x/poa/types"
	"github.com/youngjun1714/sdk/x/poa/keeper"
	"github.com/cosmos/cosmos-sdk/types/module"
)
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the poa module.
type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return types.ModuleName

type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}
