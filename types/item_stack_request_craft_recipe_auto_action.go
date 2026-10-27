package types

type ItemStackRequestCraftRecipeAutoAction struct {
	ActionType              ItemStackRequestActionType
	RecipeNetId             RecipeNetId
	NumberOfRequestedCrafts uint8
	Ingredients             []RecipeIngredient
}
