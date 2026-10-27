package types

type ItemStackRequestCraftRecipeAction struct {
	ActionType              ItemStackRequestActionType
	RecipeNetId             RecipeNetId
	NumberOfRequestedCrafts uint8
}
