package types

type AttributeModifier struct {
	Id             string
	Name           string
	Amount         float32
	Operation      int32
	Operand        int32
	IsSerializable bool
}
