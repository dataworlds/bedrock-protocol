package types

type StructureTemplateRequestOperation uint8

const (
	StructureTemplateRequestOperationNone                StructureTemplateRequestOperation = 0
	StructureTemplateRequestOperationExportFromSaveMode  StructureTemplateRequestOperation = 1
	StructureTemplateRequestOperationExportFromLoadMode  StructureTemplateRequestOperation = 2
	StructureTemplateRequestOperationQuerySavedStructure StructureTemplateRequestOperation = 3
)
