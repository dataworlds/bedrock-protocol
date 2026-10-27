package types

type TransitionAttributeData struct {
	FromAttribute struct {
		ControlValue       uint32
		BoolAttributeData  *BoolAttributeData
		FloatAttributeData *FloatAttributeData
		ColorAttributeData *ColorAttributeData
	}
	ToAttribute struct {
		ControlValue       uint32
		BoolAttributeData  *BoolAttributeData
		FloatAttributeData *FloatAttributeData
		ColorAttributeData *ColorAttributeData
	}
	Settings TransitionSettingsData
}
