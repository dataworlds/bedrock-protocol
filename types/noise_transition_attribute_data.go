package types

type NoiseTransitionAttributeData struct {
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
	Settings NoiseTransitionSettingsData
}
