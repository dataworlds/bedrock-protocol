package types

type UpdateEnvironmentAttributesData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	Attributes              []EnvironmentAttributeData
}
