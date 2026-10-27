package types

type ShapedRecipePayload struct {
	RecipeId             string
	Width                int32
	Height               int32
	Ingredients          []RecipeIngredient
	Results              []NetworkItemInstanceDescriptorData
	UUID                 MceUUID
	Tag                  string
	Priority             int32
	AssumeSymmetry       bool
	UnlockingRequirement *RecipeUnlockingRequirement
	NetId                RecipeNetId
}
