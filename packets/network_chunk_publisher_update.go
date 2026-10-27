package packets

import "github.com/dataworlds/bedrock-protocol/types"

const NetworkChunkPublisherUpdateID uint32 = 121

type NetworkChunkPublisherUpdate struct {
	NewPositionForView    types.BlockPos
	NewRadiusForView      uint32
	ServerBuiltChunksList []types.ChunkPos
}
