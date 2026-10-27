package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlayerActionID uint32 = 36

type PlayerAction struct {
	PlayerRuntimeID types.ActorRuntimeID
	Action          types.PlayerActionType
	BlockPosition   types.BlockPos
	ResultPos       types.BlockPos
	Face            int32
}
