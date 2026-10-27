package types

type DataStoreUpdate struct {
	DataStoreName string
	Property      string
	Path          string
	Data          struct {
		ControlValue uint32
		Variant0     *float64
		Variant1     *bool
		Variant2     *string
	}
	PropertyUpdateCount uint32
	PathUpdateCount     uint32
}
