package packets

import "github.com/dataworlds/bedrock-protocol/types"

const RecordStartedID uint32 = 352

type RecordStarted struct {
	BlockPosition     types.BlockPos
	ServerSoundHandle types.ServerSoundHandle
}
