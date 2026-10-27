package types

type ItemStackRequestTakeAction struct {
	ActionType  ItemStackRequestActionType
	Amount      uint8
	Source      ItemStackRequestSlotInfo
	Destination ItemStackRequestSlotInfo
}
