package packets

import "github.com/dataworlds/bedrock-protocol/types"

const InventorySlotID uint32 = 50

type InventorySlot struct {
	ContainerId       uint8
	Slot              uint32
	FullContainerName *types.FullContainerName
	StorageItem       *types.NetworkItemStackDescriptor
	Item              types.NetworkItemStackDescriptor
}
