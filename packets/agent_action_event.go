package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AgentActionEventID uint32 = 181

type AgentActionEvent struct {
	RequestId string
	Action    types.AgentActionType
	Response  types.JsonValue
}
