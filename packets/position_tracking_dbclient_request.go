package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PositionTrackingDBClientRequestID uint32 = 154

type PositionTrackingDBClientRequest struct {
	Action types.PositionTrackingDBClientRequestPacketPayloadAction
	Id     types.PositionTrackingId
}
