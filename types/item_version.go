package types

type ItemVersion int32

const (
	ItemVersionLegacy     ItemVersion = 0
	ItemVersionDataDriven ItemVersion = 1
	ItemVersionNone       ItemVersion = 2
)
