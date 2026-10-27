package types

type MoveActorAbsoluteData struct {
	ActorRuntimeID ActorRuntimeID
	Header         uint8
	Position       Vec3
	RotationX      uint8
	RotationY      uint8
	RotationYHead  uint8
}
