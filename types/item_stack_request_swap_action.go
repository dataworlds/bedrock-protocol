package types

type ItemStackRequestSwapAction struct {
	ActionType  ItemStackRequestActionType
	Source      ItemStackRequestSlotInfo
	Destination ItemStackRequestSlotInfo
}
