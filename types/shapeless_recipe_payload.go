package types

type ShapelessRecipePayload struct {
	RecipeId             string
	Ingredients          []RecipeIngredient
	Results              []NetworkItemInstanceDescriptorData
	UUID                 MceUUID
	Tag                  string
	Priority             int32
	UnlockingRequirement *RecipeUnlockingRequirement
	NetId                RecipeNetId
}
