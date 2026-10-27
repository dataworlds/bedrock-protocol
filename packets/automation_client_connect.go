package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AutomationClientConnectID uint32 = 95

type AutomationClientConnect struct {
	WebSocketData types.WebSocketPacketData
}
