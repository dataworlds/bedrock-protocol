package types

type Rotation uint8

const (
	RotationNone               Rotation = 0
	RotationRotate90           Rotation = 1
	RotationRotate180          Rotation = 2
	RotationRotate270          Rotation = 3
	RotationClockwise90        Rotation = 1
	RotationClockwise180       Rotation = 2
	RotationCounterClockwise90 Rotation = 3
)
