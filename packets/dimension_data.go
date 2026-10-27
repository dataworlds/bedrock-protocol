package packets

import "github.com/dataworlds/bedrock-protocol/types"

const DimensionDataID uint32 = 180

type DimensionData struct {
	Definitions map[string]types.DimensionDefinition
}
