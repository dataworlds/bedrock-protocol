package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MoveActorAbsoluteID uint32 = 18

type MoveActorAbsolute struct {
	MoveData types.MoveActorAbsoluteData
}
