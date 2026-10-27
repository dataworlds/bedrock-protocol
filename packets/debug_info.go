package packets

import "github.com/dataworlds/bedrock-protocol/types"

const DebugInfoID uint32 = 155

type DebugInfo struct {
	ActorId types.ActorUniqueID
	Data    string
}
