package packets

const ResourcePackChunkDataID uint32 = 83

type ResourcePackChunkData struct {
	ResourceName string
	ChunkID      uint32
	ByteOffset   uint64
	ChunkData    string
}
