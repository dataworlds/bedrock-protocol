package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdateAttributesID uint32 = 29

type UpdateAttributes struct {
	TargetRuntimeID types.ActorRuntimeID
	AttributeList   []types.AttributeData
	Tick            types.PlayerInputTick
}
