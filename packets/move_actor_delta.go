package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MoveActorDeltaID uint32 = 111

type MoveActorDelta struct {
	MoveData types.MoveActorDeltaData
}
