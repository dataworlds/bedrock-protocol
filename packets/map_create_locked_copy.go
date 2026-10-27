package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MapCreateLockedCopyID uint32 = 131

type MapCreateLockedCopy struct {
	OriginalMapId types.ActorUniqueID
	NewMapId      types.ActorUniqueID
}
