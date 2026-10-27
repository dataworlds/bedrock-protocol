package types

type BiomeReplacementData struct {
	ReplacementBiome    uint16
	Dimension           uint16
	TargetBiomes        []uint16
	Amount              float32
	NoiseFrequencyScale float32
	ReplacementIndex    uint32
}
