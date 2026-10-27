package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetSpawnPositionID uint32 = 43

type SetSpawnPosition struct {
	SpawnPositionType types.SpawnPositionType
	BlockPosition     types.BlockPos
	DimensionType     types.DimensionType
	SpawnBlockPos     types.BlockPos
}
