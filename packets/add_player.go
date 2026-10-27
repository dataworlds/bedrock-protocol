package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AddPlayerID uint32 = 12

type AddPlayer struct {
	UUID               types.MceUUID
	PlayerName         string
	TargetRuntimeID    types.ActorRuntimeID
	PlatformChatId     string
	Position           types.Vec3
	Velocity           types.Vec3
	Rotation           types.Vec2
	YHeadRotation      float32
	CarriedItem        types.NetworkItemStackDescriptor
	PlayerGameType     types.GameType
	EntityData         types.SynchedActorDataList
	SynchedProperties  types.PropertySyncData
	AbilitiesData      types.SerializedAbilitiesData
	ActorLinks         []types.ActorLink
	PassengerBlockData *types.PassengerOfBlockArguments
	DeviceId           string
	BuildPlatform      types.BuildPlatform
}
