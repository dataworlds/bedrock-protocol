package types

type SubChunkPacketData struct {
	SubChunkPosOffset     SubChunkPosOffset
	SubChunkRequestResult uint8
	SerializedSubChunk    *string
	HeightMapData         SubChunkHeightmapData
	BlobId                *uint64
}
