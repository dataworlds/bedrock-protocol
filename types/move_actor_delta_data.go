package types

type MoveActorDeltaData struct {
	ActorRuntimeID       ActorRuntimeID
	NewPositionX         *float32
	NewPositionY         *float32
	NewPositionZ         *float32
	RotationX            *int8
	RotationY            *int8
	RotationYHead        *int8
	IsOnGround           bool
	ForceMove            bool
	ForceMoveLocalEntity bool
	ForceCompletion      bool
	Ticks                uint64
}
