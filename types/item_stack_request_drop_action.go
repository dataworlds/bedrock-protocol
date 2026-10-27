package types

type ItemStackRequestDropAction struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     ItemStackRequestSlotInfo
	Randomly   bool
}
