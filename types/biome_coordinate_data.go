package types

type BiomeCoordinateData struct {
	MinValueType int32
	MinValue     uint16
	MaxValueType int32
	MaxValue     uint16
	GridOffset   uint32
	GridStepSize uint32
	Distribution RandomDistributionType
}
