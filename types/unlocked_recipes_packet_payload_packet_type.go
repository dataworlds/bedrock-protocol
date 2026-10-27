package types

type UnlockedRecipesPacketPayloadPacketType uint32

const (
	UnlockedRecipesPacketPayloadPacketTypeEmpty                    UnlockedRecipesPacketPayloadPacketType = 0
	UnlockedRecipesPacketPayloadPacketTypeInitiallyUnlockedRecipes UnlockedRecipesPacketPayloadPacketType = 1
	UnlockedRecipesPacketPayloadPacketTypeNewlyUnlockedRecipes     UnlockedRecipesPacketPayloadPacketType = 2
	UnlockedRecipesPacketPayloadPacketTypeRemoveUnlockedRecipes    UnlockedRecipesPacketPayloadPacketType = 3
	UnlockedRecipesPacketPayloadPacketTypeRemoveAllUnlockedRecipes UnlockedRecipesPacketPayloadPacketType = 4
)
