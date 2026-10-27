package types

type BiomeElementData struct {
	NoiseFreqScale    float32
	NoiseLowerBound   float32
	NoiseUpperBound   float32
	HeightMinType     int32
	HeightMin         uint16
	HeightMaxType     int32
	HeightMax         uint16
	AdjustedMaterials BiomeSurfaceMaterialData
}
