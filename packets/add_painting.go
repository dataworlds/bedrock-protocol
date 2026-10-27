package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AddPaintingID uint32 = 22

type AddPainting struct {
	TargetActorID   types.ActorUniqueID
	TargetRuntimeID types.ActorRuntimeID
	Position        types.Vec3
	Direction       int32
	Motif           string
}
