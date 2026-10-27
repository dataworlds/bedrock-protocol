package packets

const RequestChunkRadiusID uint32 = 69

type RequestChunkRadius struct {
	ChunkRadius    int32
	MaxChunkRadius uint8
}
