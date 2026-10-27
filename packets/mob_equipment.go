package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MobEquipmentID uint32 = 31

type MobEquipment struct {
	TargetRuntimeID types.ActorRuntimeID
	Item            types.NetworkItemStackDescriptor
	Slot            uint8
	SelectedSlot    uint8
	ContainerID     uint8
}
