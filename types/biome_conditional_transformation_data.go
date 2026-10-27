package types

type BiomeConditionalTransformationData struct {
	TransformsInto      []BiomeWeightedData
	ConditionJson       uint16
	MinPassingNeighbors uint32
}
