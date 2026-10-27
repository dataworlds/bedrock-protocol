package packets

import "github.com/dataworlds/bedrock-protocol/types"

const InventoryContentID uint32 = 49

type InventoryContent struct {
	ContainerId       uint32
	Slots             []types.NetworkItemStackDescriptor
	FullContainerName types.FullContainerName
	StorageItem       types.NetworkItemStackDescriptor
}
