package packets

import "github.com/dataworlds/bedrock-protocol/types"

const StructureBlockUpdateID uint32 = 90

type StructureBlockUpdate struct {
	BlockPosition types.BlockPos
	StructureData types.StructureEditorData
	Trigger       bool
	IsWaterlogged bool
}
