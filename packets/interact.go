package packets

import "github.com/dataworlds/bedrock-protocol/types"

const InteractID uint32 = 33

type Interact struct {
	Action          types.InteractPacketPayloadAction
	TargetRuntimeID types.ActorRuntimeID
	Position        *types.Vec3
}
