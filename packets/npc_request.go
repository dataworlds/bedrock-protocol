package packets

import "github.com/dataworlds/bedrock-protocol/types"

const NpcRequestID uint32 = 98

type NpcRequest struct {
	NPCRuntimeID types.ActorRuntimeID
	RequestType  types.NpcRequestPacketPayloadRequestType
	Actions      string
	ActionIndex  uint8
	SceneName    string
}
