package packets

import "github.com/dataworlds/bedrock-protocol/types"

const LecternUpdateID uint32 = 125

type LecternUpdate struct {
	NewPageToShow             uint8
	TotalPages                uint8
	PositionOfLecternToUpdate types.BlockPos
}
