package types

type BiomeTemperatureCategory uint8

const (
	BiomeTemperatureCategoryMedium   BiomeTemperatureCategory = 0
	BiomeTemperatureCategoryWarm     BiomeTemperatureCategory = 1
	BiomeTemperatureCategoryLukewarm BiomeTemperatureCategory = 2
	BiomeTemperatureCategoryCold     BiomeTemperatureCategory = 3
	BiomeTemperatureCategoryFrozen   BiomeTemperatureCategory = 4
)
