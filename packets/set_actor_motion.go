package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetActorMotionID uint32 = 40

type SetActorMotion struct {
	TargetRuntimeID types.ActorRuntimeID
	Motion          types.Vec3
	Tick            types.PlayerInputTick
}
