package types

type BiomeScatterParamData struct {
	Coordinates       []BiomeCoordinateData
	EvalOrder         CoordinateEvaluationOrder
	ChancePercentType int32
	ChancePercent     uint16
	ChanceNumerator   int32
	ChanceDenominator int32
	IterationsType    int32
	Iterations        uint16
}
