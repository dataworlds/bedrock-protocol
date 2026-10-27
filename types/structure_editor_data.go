package types

type StructureEditorData struct {
	StructureName         BedrockSafetyRedactableString
	DataField             string
	ShouldIncludePlayers  bool
	ShouldShowBoundingBox bool
	StructureBlockType    StructureBlockType
	StructureSettings     StructureSettings
	RedstoneSaveMode      StructureRedstoneSaveMode
}
