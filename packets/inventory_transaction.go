package packets

import "github.com/dataworlds/bedrock-protocol/types"

const InventoryTransactionID uint32 = 30

type InventoryTransaction struct {
	LegacyRequestID    types.ItemStackLegacyRequestId
	LegacySetItemSlots *[]types.LegacySetSlot
	Transaction        struct {
		ControlValue                       uint32
		NormalTransactionData              *types.NormalTransactionData
		InventoryMismatchData              *types.InventoryMismatchData
		ItemUseInventoryTransaction        *types.ItemUseInventoryTransaction
		ItemUseOnActorInventoryTransaction *types.ItemUseOnActorInventoryTransaction
		ItemReleaseInventoryTransaction    *types.ItemReleaseInventoryTransaction
	}
}
