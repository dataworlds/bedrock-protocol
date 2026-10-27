package types

type SmithingTransformRecipePayload struct {
	RecipeId           string
	TemplateIngredient RecipeIngredient
	BaseIngredient     RecipeIngredient
	AdditionIngredient RecipeIngredient
	Result             NetworkItemInstanceDescriptorData
	Tag                string
	NetId              RecipeNetId
}
