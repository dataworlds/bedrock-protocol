package types

type ItemStackRequestDestroyAction struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     ItemStackRequestSlotInfo
}
