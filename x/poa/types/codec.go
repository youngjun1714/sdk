package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgSubmitApplication{}, "poa/MsgSubmitApplication", nil)
	cdc.RegisterConcrete(MsgVote{}, "poa/MsgVote", nil)
	cdc.RegisterConcrete(MsgProposeKick{}, "poa/MsgProposeKick", nil)
	cdc.RegisterConcrete(MsgLeaveValidatorSet{}, "poa/MsgLeaveValidatorSet", nil)
}

// ModuleCdc defines the module codec
var (
	amino = codec.NewLegacyAmino()

	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

/*var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.NewAminoCodec()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}*/
