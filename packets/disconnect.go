package packets

import "github.com/dataworlds/bedrock-protocol/types"

const DisconnectID uint32 = 5

type Disconnect struct {
	Reason   types.DisconnectFailReason
	Messages struct {
		ControlValue             uint32
		DisconnectPacketMessages *types.DisconnectPacketMessages
		Variant1                 *any
	}
}
