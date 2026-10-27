package packets

import "github.com/dataworlds/bedrock-protocol/types"

const RemoveActorID uint32 = 14

type RemoveActor struct {
	TargetActorID types.ActorUniqueID
}
