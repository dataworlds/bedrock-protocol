package types

type PlayerBlockActionData struct {
	PlayerActionType PlayerActionType
	Position         BlockPos
	Facing           int32
}
