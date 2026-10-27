package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PrimitiveShapesID uint32 = 328

type PrimitiveShapes struct {
	ArrayOfPrimitiveShapesCanBeAMixOfNewUpdatedOrRemoved []types.PrimitiveShapeDataPayload
}
