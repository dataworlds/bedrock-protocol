package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UnlockedRecipesID uint32 = 199

type UnlockedRecipes struct {
	PacketType          types.UnlockedRecipesPacketPayloadPacketType
	UnlockedRecipesList []string
}
