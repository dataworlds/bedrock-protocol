package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SubChunkRequestID uint32 = 175

type SubChunkRequest struct {
	DimensionType              types.DimensionType
	SubChunkPositionOffsetList []types.SubChunkPosOffset
	CenterPos                  types.SubChunkPos
}
