package packets

import "github.com/dataworlds/bedrock-protocol/types"

const TakeItemActorID uint32 = 17

type TakeItemActor struct {
	ItemRuntimeID  types.ActorRuntimeID
	ActorRuntimeID types.ActorRuntimeID
}
