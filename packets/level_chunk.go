package packets

import "github.com/dataworlds/bedrock-protocol/types"

const LevelChunkID uint32 = 58

type LevelChunk struct {
	ChunkPosition              types.ChunkPos
	DimensionId                types.DimensionType
	SubChunksCount             uint32
	ClientRequestSubChunkLimit *int32
	CacheEnabled               bool
	CacheMetadata              []types.LevelChunkPacketPayloadSubChunkMetadata
	SerializedChunkData        string
	IsClientBiomeUpdate        bool
}
