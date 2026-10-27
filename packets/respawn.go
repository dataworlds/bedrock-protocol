package packets

import "github.com/dataworlds/bedrock-protocol/types"

const RespawnID uint32 = 45

type Respawn struct {
	Position        types.Vec3
	State           types.PlayerRespawnState
	PlayerRuntimeId types.ActorRuntimeID
}
