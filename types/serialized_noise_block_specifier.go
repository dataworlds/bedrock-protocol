package types

type SerializedNoiseBlockSpecifier struct {
	Noise     *string
	Threshold *float32
	Range     *FloatRange
	Block     *uint32
}
