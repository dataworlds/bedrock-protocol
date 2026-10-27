package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientboundDebugRendererID uint32 = 164

type ClientboundDebugRenderer struct {
	Type            uint8
	DebugMarkerData *types.DebugMarkerData
}
