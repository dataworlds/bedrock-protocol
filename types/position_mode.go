package types

type PositionMode uint8

const (
	PositionModeNormal      PositionMode = 0
	PositionModeRespawn     PositionMode = 1
	PositionModeTeleport    PositionMode = 2
	PositionModeOnlyHeadRot PositionMode = 3
)
