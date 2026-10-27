package types

type CameraAimAssistPresetExclusionDefinition struct {
	Blocks             *[]BlockReference
	Entities           *[]EntityReference
	BlockTags          *[]string
	EntityTypeFamilies *[]string
}
