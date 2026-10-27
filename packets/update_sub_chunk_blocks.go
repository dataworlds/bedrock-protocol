package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdateSubChunkBlocksID uint32 = 172

type UpdateSubChunkBlocks struct {
	SubChunkBlockPosition types.BlockPos
	BlocksChanged         types.UpdateSubChunkBlocksChangedInfo
}
