package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientboundStonecutterSetRecipeID uint32 = 355

type ClientboundStonecutterSetRecipe struct {
	PlayerId    types.ActorUniqueID
	ContainerId uint8
	RecipeIndex int32
}
