package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientboundTextureShiftID uint32 = 336

type ClientboundTextureShift struct {
	ActionID             types.ClientboundTextureShiftPacketPayloadAction
	CollectionName       string
	FromStep             string
	ToStep               string
	AllSteps             []string
	CurrentLengthInTicks uint64
	TotalLengthInTicks   uint64
	Enabled              bool
}
