package types

type InventoryLayout int32

const (
	InventoryLayoutNone           InventoryLayout = 0
	InventoryLayoutInventoryOnly  InventoryLayout = 1
	InventoryLayoutDefault        InventoryLayout = 2
	InventoryLayoutRecipeBookOnly InventoryLayout = 3
)
