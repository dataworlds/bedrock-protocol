package types

type ItemReleaseInventoryTransaction struct {
	Actions      InventoryTransaction
	ActionType   ItemReleaseActionType
	Slot         int32
	Item         NetworkItemStackDescriptor
	FromPosition Vec3
	Hand         HandSlot
}
