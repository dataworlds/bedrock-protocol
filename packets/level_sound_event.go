package packets

import "github.com/dataworlds/bedrock-protocol/types"

const LevelSoundEventID uint32 = 123

type LevelSoundEvent struct {
	SoundEvent      types.SoundEventIdentifier
	Position        types.Vec3
	Data            int32
	ActorIdentifier string
	IsBaby          bool
	IsGlobal        bool
	ActorUniqueId   int64
	FireAtPosition  *types.Vec3
}
