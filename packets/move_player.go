package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MovePlayerID uint32 = 19

type MovePlayer struct {
	PlayerRuntimeID types.ActorRuntimeID
	Position        types.Vec3
	Rotation        types.Vec2
	YHeadRotation   float32
	PositionMode    uint8
	OnGround        bool
	RidingRuntimeID types.ActorRuntimeID
	TeleportData    *types.MovePlayerTeleportData
	Tick            types.PlayerInputTick
}
