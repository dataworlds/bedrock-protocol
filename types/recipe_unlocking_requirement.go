package types

type RecipeUnlockingRequirement struct {
	UnlockingContext     RecipeUnlockingContext
	UnlockingIngredients *[]RecipeIngredient
}
