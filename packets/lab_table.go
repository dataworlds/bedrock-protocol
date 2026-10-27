package packets

import "github.com/dataworlds/bedrock-protocol/types"

const LabTableID uint32 = 109

type LabTable struct {
	Type     types.LabTablePacketPayloadType
	Position types.BlockPos
	Reaction types.LabTableReactionType
}
