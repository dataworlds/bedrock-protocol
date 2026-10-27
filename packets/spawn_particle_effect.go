package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SpawnParticleEffectID uint32 = 118

type SpawnParticleEffect struct {
	DimensionId     uint8
	ActorId         types.ActorUniqueID
	Position        types.Vec3
	EffectName      string
	MolangVariables *types.MolangVariableMap
}
