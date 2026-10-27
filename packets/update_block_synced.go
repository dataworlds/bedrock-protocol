package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdateBlockSyncedID uint32 = 110

type UpdateBlockSynced struct {
	BlockPosition    types.BlockPos
	BlockRuntimeID   uint32
	Flags            uint32
	Layer            uint32
	UniqueActorId    uint64
	ActorSyncMessage uint64
}
