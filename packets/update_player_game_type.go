package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdatePlayerGameTypeID uint32 = 151

type UpdatePlayerGameType struct {
	PlayerGameType types.GameType
	TargetPlayer   types.ActorUniqueID
	Tick           types.PlayerInputTick
}
