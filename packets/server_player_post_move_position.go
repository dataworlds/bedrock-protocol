package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ServerPlayerPostMovePositionID uint32 = 16

type ServerPlayerPostMovePosition struct {
	Pos types.Vec3
}
