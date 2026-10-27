package types

type InventorySource struct {
	SourceType  InventorySourceType
	ContainerID *int8
	BitFlags    *InventorySourceFlags
}
