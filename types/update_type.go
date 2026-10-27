package types

type UpdateType uint8

const (
	UpdateTypeClearOverrides   UpdateType = 0
	UpdateTypeRemoveOverride   UpdateType = 1
	UpdateTypeSetIntOverride   UpdateType = 2
	UpdateTypeSetFloatOverride UpdateType = 3
)
