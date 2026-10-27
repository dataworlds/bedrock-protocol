package types

type NewInteractionModel int32

const (
	NewInteractionModelTouch     NewInteractionModel = 0
	NewInteractionModelCrosshair NewInteractionModel = 1
	NewInteractionModelClassic   NewInteractionModel = 2
	NewInteractionModelCount     NewInteractionModel = 3
)
