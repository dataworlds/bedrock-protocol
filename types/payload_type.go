package types

type PayloadType uint8

const (
	PayloadTypeInvalid            PayloadType = 0
	PayloadTypeClearDebugMarkers  PayloadType = 1
	PayloadTypeAddDebugMarkerCube PayloadType = 2
)
