package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AnimateID uint32 = 44

type Animate struct {
	Action               types.AnimatePacketPayloadAction
	TargetActorRuntimeID types.ActorRuntimeID
	Data                 float32
	SwingSource          *types.ActorSwingSource
	Hand                 *types.HandSlot
}
