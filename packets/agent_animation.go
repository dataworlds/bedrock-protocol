package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AgentAnimationID uint32 = 304

type AgentAnimation struct {
	AgentAnimation types.AgentAnimation
	RuntimeId      types.ActorRuntimeID
}
