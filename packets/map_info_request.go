package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MapInfoRequestID uint32 = 68

type MapInfoRequest struct {
	MapUniqueID      types.ActorUniqueID
	ClientPixelsList []types.AnonymousNamespaceMapInfoRequestPacketAnonClientPixelsProxy
}
