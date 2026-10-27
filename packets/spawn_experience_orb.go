package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SpawnExperienceOrbID uint32 = 66

type SpawnExperienceOrb struct {
	Position types.Vec3
	XPValue  int32
}
