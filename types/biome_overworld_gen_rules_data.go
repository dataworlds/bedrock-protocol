package types

type BiomeOverworldGenRulesData struct {
	HillsTransformations  []BiomeWeightedData
	MutateTransformations []BiomeWeightedData
	RiverTransformations  []BiomeWeightedData
	ShoreTransformations  []BiomeWeightedData
	PreHillsEdge          []BiomeConditionalTransformationData
	PostShoreEdge         []BiomeConditionalTransformationData
	Climate               []BiomeWeightedTemperatureData
}
