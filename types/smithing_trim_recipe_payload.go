package types

type SmithingTrimRecipePayload struct {
	RecipeId           string
	TemplateIngredient RecipeIngredient
	BaseIngredient     RecipeIngredient
	AdditionIngredient RecipeIngredient
	Tag                string
	NetId              RecipeNetId
}
