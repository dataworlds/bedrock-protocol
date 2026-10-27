package types

type BiomeConsolidatedFeatureData struct {
	Scatter               BiomeScatterParamData
	Feature               uint16
	Identifier            uint16
	Pass                  uint16
	CanUseInternalFeature bool
}
