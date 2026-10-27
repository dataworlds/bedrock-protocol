package types

type ItemStackRequestCraftRepairAndDisenchantAction struct {
	ActionType              ItemStackRequestActionType
	RecipeNetId             ItemStackNetIdVariant
	NumberOfRequestedCrafts uint8
	RepairCost              int32
}
