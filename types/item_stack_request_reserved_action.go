package types

type ItemStackRequestReservedAction struct {
	ActionType ItemStackRequestActionType
	ReservedId string
	NumCrafts  uint8
}
