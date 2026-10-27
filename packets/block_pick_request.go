package packets

import "github.com/dataworlds/bedrock-protocol/types"

const BlockPickRequestID uint32 = 34

type BlockPickRequest struct {
	Position types.BlockPos
	WithData bool
	MaxSlots uint8
}
