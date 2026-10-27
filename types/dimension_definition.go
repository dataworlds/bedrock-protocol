package types

type DimensionDefinition struct {
	MinimumY      int32
	HeightRange   int32
	GeneratorType int32
	DimensionType DimensionType
	PackId        MceUUID
	DefaultBiome  string
	CloudHeight   int32
	RenderClouds  bool
}
