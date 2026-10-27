package types

type ItemStackRequestCraftLoomAction struct {
	ActionType    ItemStackRequestActionType
	PatternNameId string
	NumCrafts     uint8
}
