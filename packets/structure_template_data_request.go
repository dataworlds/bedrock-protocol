package packets

import "github.com/dataworlds/bedrock-protocol/types"

const StructureTemplateDataRequestID uint32 = 132

type StructureTemplateDataRequest struct {
	StructureName      string
	StructurePosition  types.BlockPos
	StructureSettings  types.StructureSettings
	RequestedOperation types.StructureTemplateRequestOperation
}
