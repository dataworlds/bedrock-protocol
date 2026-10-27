package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MobArmorEquipmentID uint32 = 32

type MobArmorEquipment struct {
	TargetRuntimeID types.ActorRuntimeID
	Head            types.NetworkItemStackDescriptor
	Torso           types.NetworkItemStackDescriptor
	Legs            types.NetworkItemStackDescriptor
	Feet            types.NetworkItemStackDescriptor
	Body            types.NetworkItemStackDescriptor
}
