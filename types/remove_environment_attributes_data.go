package types

type RemoveEnvironmentAttributesData struct {
	AttributeLayerName      string
	AttributeLayerDimension DimensionType
	Attributes              []string
}
