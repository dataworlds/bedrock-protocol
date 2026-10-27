package packets

const EditorNetworkID uint32 = 190

type EditorNetwork struct {
	RouteToManager bool
	RawVariantName string
	RawVariantData string
}
