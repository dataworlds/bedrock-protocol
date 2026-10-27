package packets

import "github.com/dataworlds/bedrock-protocol/types"

const LevelEventID uint32 = 25

type LevelEvent struct {
	EventId  int32
	Position types.Vec3
	Data     int32
}
