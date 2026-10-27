package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PositionTrackingDBServerBroadcastID uint32 = 153

type PositionTrackingDBServerBroadcast struct {
	Action               types.PositionTrackingDBServerBroadcastPacketPayloadAction
	Id                   types.PositionTrackingId
	PositionTrackingData any
}
