package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ChangeDimensionID uint32 = 61

type ChangeDimension struct {
	DimensionID     types.DimensionType
	Position        types.Vec3
	Respawn         bool
	LoadingScreenId *uint32
}
