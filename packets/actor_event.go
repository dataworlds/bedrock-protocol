package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ActorEventID uint32 = 27

type ActorEvent struct {
	TargetRuntimeID types.ActorRuntimeID
	EventID         types.ActorEvent
	Data            int32
	FireAtPosition  *types.Vec3
}
