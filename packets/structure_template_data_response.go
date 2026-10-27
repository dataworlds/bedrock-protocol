package packets

import "github.com/dataworlds/bedrock-protocol/types"

const StructureTemplateDataResponseID uint32 = 133

type StructureTemplateDataResponse struct {
	StructureName string
	StructureSNBT any
	ResponseType  types.StructureTemplateResponseType
}
