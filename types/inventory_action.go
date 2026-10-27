package types

type InventoryAction struct {
	Source   InventorySource
	Slot     uint32
	FromItem NetworkItemStackDescriptor
	ToItem   NetworkItemStackDescriptor
}
