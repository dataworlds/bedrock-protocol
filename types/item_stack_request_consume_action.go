package types

type ItemStackRequestConsumeAction struct {
	ActionType ItemStackRequestActionType
	Amount     uint8
	Source     ItemStackRequestSlotInfo
}
