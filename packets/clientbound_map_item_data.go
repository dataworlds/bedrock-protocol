package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientboundMapItemDataID uint32 = 67

type ClientboundMapItemData struct {
	MapID           types.ActorUniqueID
	Dimension       uint8
	IsLocked        bool
	MapOrigin       types.BlockPos
	CreationMapIDs  *[]types.ActorUniqueID
	Scale           *int8
	TrackedActorIDs *[]types.MapItemTrackedActorUniqueId
	Decorations     *[]types.MapDecoration
	Width           *int32
	Height          *int32
	StartX          *int32
	StartY          *int32
	Pixels          *[]uint32
}
