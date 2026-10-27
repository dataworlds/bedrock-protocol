package packets

import "github.com/dataworlds/bedrock-protocol/types"

const BlockEventID uint32 = 26

type BlockEvent struct {
	BlockPosition types.BlockPos
	EventType     int32
	EventValue    int32
}
