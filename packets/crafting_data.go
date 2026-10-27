package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CraftingDataID uint32 = 52

type CraftingData struct {
	ShapedRecipes             []types.ShapedRecipePayload
	ShapelessRecipes          []types.ShapelessRecipePayload
	MultiRecipes              []types.MultiRecipePayload
	UserDataShapelessRecipes  []types.ShapelessRecipePayload
	ShapelessChemistryRecipes []types.ShapelessRecipePayload
	ShapedChemistryRecipes    []types.ShapedRecipePayload
	SmithingTransformRecipes  []types.SmithingTransformRecipePayload
	SmithingTrimRecipes       []types.SmithingTrimRecipePayload
	PotionMixes               []types.PotionMixDataEntry
	ContainerMixes            []types.ContainerMixDataEntry
	MaterialReducers          []types.MaterialReducerDataEntry
	ClearRecipes              bool
}
