package types

type BiomeDefinitionData struct {
	Id                uint16
	Temperature       float32
	Downfall          float32
	FoliageSnow       float32
	Depth             float32
	Scale             float32
	MapWaterColorARGB int32
	Rain              bool
	Tags              *BiomeTagsData
	ChunkGenData      *BiomeDefinitionChunkGenData
}
