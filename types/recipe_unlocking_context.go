package types

type RecipeUnlockingContext int32

const (
	RecipeUnlockingContextNone               RecipeUnlockingContext = 0
	RecipeUnlockingContextAlwaysUnlocked     RecipeUnlockingContext = 1
	RecipeUnlockingContextPlayerInWater      RecipeUnlockingContext = 2
	RecipeUnlockingContextPlayerHasManyItems RecipeUnlockingContext = 3
)
