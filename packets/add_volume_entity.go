package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AddVolumeEntityID uint32 = 166

type AddVolumeEntity struct {
	EntityNetworkId types.EntityNetId
	Components      any
	JSONIdentifier  string
	InstanceName    string
	MinBounds       types.BlockPos
	MaxBounds       types.BlockPos
	DimensionType   types.DimensionType
	EngineVersion   string
}
