package types

type CameraAimAssistPresetDefinition struct {
	Identifier          CameraAimAssistIdentifier
	ExclusionSettings   *CameraAimAssistPresetExclusionDefinition
	LiquidTargetingList *[]ItemReference
	ItemSettings        *map[ItemReference]CameraAimAssistCategoryReference
	DefaultItemSettings *CameraAimAssistCategoryReference
	HandSettings        *CameraAimAssistCategoryReference
}
