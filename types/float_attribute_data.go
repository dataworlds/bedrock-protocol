package types

type FloatAttributeData struct {
	Value         float32
	Operation     *FloatAttributeOperation
	ConstraintMin *float32
	ConstraintMax *float32
}
