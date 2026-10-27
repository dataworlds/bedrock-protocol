package types

type CameraAimAssistCategoryPriorities struct {
	Entities           *map[EntityReference]int32
	Blocks             *map[BlockReference]int32
	BlockTags          *map[string]int32
	EntityTypeFamilies *map[string]int32
	EntityDefault      *int32
	BlockDefault       *int32
}
