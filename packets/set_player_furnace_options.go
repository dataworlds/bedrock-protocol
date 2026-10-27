package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetPlayerFurnaceOptionsID uint32 = 351

type SetPlayerFurnaceOptions struct {
	FurnaceType    types.SetPlayerFurnaceOptionsPacketPayloadFurnaceType
	FurnaceOptions types.FurnaceOptions
}
