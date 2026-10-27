package types

type ItemStackRequestPlaceAction struct {
	ActionType  ItemStackRequestActionType
	Amount      uint8
	Source      ItemStackRequestSlotInfo
	Destination ItemStackRequestSlotInfo
}
