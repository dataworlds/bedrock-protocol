package packets

const ResourcePackDataInfoID uint32 = 82

type ResourcePackDataInfo struct {
	ResourceName   string
	ChunkSize      uint32
	NumberOfChunks uint32
	FileSize       uint64
	FileHash       string
	IsPremiumPack  bool
	PackType       uint8
}
