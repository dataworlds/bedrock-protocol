package types

type AttributeLayerData struct {
	Name       string
	Dimension  DimensionType
	Settings   AttributeLayerSettings
	Attributes []EnvironmentAttributeData
}
