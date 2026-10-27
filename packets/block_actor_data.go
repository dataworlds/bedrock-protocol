package packets

import "github.com/dataworlds/bedrock-protocol/types"

const BlockActorDataID uint32 = 56

type BlockActorData struct {
	BlockPosition types.BlockPos
	ActorDataTags any
}
