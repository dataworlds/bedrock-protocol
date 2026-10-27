package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AddActorID uint32 = 13

type AddActor struct {
	TargetActorID      types.ActorUniqueID
	TargetRuntimeID    types.ActorRuntimeID
	ActorType          string
	Position           types.Vec3
	Velocity           types.Vec3
	Rotation           types.Vec2
	YHeadRotation      float32
	YBodyRotation      float32
	AttributesList     []types.SyncedAttribute
	ActorData          types.SynchedActorDataList
	SynchedProperties  types.PropertySyncData
	ActorLinks         []types.ActorLink
	PassengerBlockData *types.PassengerOfBlockArguments
}
