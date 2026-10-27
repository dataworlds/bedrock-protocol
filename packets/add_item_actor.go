package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AddItemActorID uint32 = 15

type AddItemActor struct {
	TargetActorID   types.ActorUniqueID
	TargetRuntimeID types.ActorRuntimeID
	Item            types.NetworkItemStackDescriptor
	Position        types.Vec3
	Velocity        types.Vec3
	EntityData      types.SynchedActorDataList
	IsFromFishing   bool
}
