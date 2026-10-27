package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SubChunkID uint32 = 174

type SubChunk struct {
	CacheEnabled  bool
	DimensionType types.DimensionType
	CenterPos     types.SubChunkPos
	SubChunkData  []types.SubChunkPacketData
}
