package packets

import "github.com/dataworlds/bedrock-protocol/types"

const VoxelShapesID uint32 = 337

type VoxelShapes struct {
	Shapes           []types.VoxelShapesSerializableVoxelShape
	NameMap          map[types.HashedString]types.VoxelShapesRegistryHandle
	CustomShapeCount uint16
}
