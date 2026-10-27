package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MovementEffectID uint32 = 318

type MovementEffect struct {
	TargetRuntimeID types.ActorRuntimeID
	EffectID        types.MovementEffectType
	EffectDuration  int32
	Tick            types.PlayerInputTick
}
