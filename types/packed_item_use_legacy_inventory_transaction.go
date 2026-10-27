package types

type PackedItemUseLegacyInventoryTransaction struct {
	LegacyRequestID    ItemStackLegacyRequestId
	LegacySetItemSlots *[]LegacySetSlot
	ItemUseTransaction ItemUseInventoryTransaction
}
