package packets

import "github.com/dataworlds/bedrock-protocol/types"

const RemoveVolumeEntityID uint32 = 167

type RemoveVolumeEntity struct {
	EntityNetworkId types.EntityNetId
	DimensionType   types.DimensionType
}
