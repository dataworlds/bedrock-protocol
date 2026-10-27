package types

type SoftEnumUpdateType uint8

const (
	SoftEnumUpdateTypeAdd     SoftEnumUpdateType = 0
	SoftEnumUpdateTypeRemove  SoftEnumUpdateType = 1
	SoftEnumUpdateTypeReplace SoftEnumUpdateType = 2
)
