package packets

const ResourcePackChunkRequestID uint32 = 84

type ResourcePackChunkRequest struct {
	ResourceName string
	Chunk        int32
}
