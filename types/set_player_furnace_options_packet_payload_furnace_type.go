package types

type SetPlayerFurnaceOptionsPacketPayloadFurnaceType uint8

const (
	SetPlayerFurnaceOptionsPacketPayloadFurnaceTypeNone         SetPlayerFurnaceOptionsPacketPayloadFurnaceType = 0
	SetPlayerFurnaceOptionsPacketPayloadFurnaceTypeFurnace      SetPlayerFurnaceOptionsPacketPayloadFurnaceType = 1
	SetPlayerFurnaceOptionsPacketPayloadFurnaceTypeBlastFurnace SetPlayerFurnaceOptionsPacketPayloadFurnaceType = 2
	SetPlayerFurnaceOptionsPacketPayloadFurnaceTypeSmoker       SetPlayerFurnaceOptionsPacketPayloadFurnaceType = 3
)
