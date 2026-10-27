package types

type ConstantAttributeData struct {
	Attribute struct {
		ControlValue       uint32
		BoolAttributeData  *BoolAttributeData
		FloatAttributeData *FloatAttributeData
		ColorAttributeData *ColorAttributeData
	}
}
