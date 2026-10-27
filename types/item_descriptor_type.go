package types

type ItemDescriptorType uint8

const (
	ItemDescriptorTypeEmpty    ItemDescriptorType = 0
	ItemDescriptorTypeItemName ItemDescriptorType = 1
	ItemDescriptorTypeMolang   ItemDescriptorType = 2
	ItemDescriptorTypeItemTag  ItemDescriptorType = 3
)
