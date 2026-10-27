package types

type ItemStackRequestCraftCreativeAction struct {
	ActionType              ItemStackRequestActionType
	CreativeItemNetId       uint32
	NumberOfRequestedCrafts uint8
}
