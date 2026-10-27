package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlaySoundID uint32 = 86

type PlaySound struct {
	Name                     string
	Position                 types.BlockPos
	Volume                   float32
	Pitch                    float32
	LoopCount                int32
	BypassListenerRangeCheck bool
	ServerSoundHandle        *types.ServerSoundHandle
	PlaybackPositionSeconds  *float32
}
