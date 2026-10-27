package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MobEffectID uint32 = 28

type MobEffect struct {
	TargetRuntimeID     types.ActorRuntimeID
	EventID             types.MobEffectPacketPayloadEvent
	EffectID            int32
	EffectAmplifier     int32
	ShowParticles       bool
	EffectDurationTicks int32
	Tick                types.PlayerInputTick
	Ambient             bool
}
