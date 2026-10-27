package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdateBlockID uint32 = 21

type UpdateBlock struct {
	BlockPosition  types.BlockPos
	BlockRuntimeID uint32
	Flags          uint32
	Layer          uint32
}
