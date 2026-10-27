package packets

import "github.com/dataworlds/bedrock-protocol/types"

const OpenSignID uint32 = 303

type OpenSign struct {
	Pos         types.BlockPos
	IsFrontSide bool
}
