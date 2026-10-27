package types

type EnvironmentAttributeData struct {
	AttributeName string
	Payload       struct {
		ControlValue                 uint32
		ConstantAttributeData        *ConstantAttributeData
		TransitionAttributeData      *TransitionAttributeData
		NoiseTransitionAttributeData *NoiseTransitionAttributeData
	}
}
