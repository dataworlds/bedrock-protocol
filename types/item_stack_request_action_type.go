package types

type ItemStackRequestActionType uint8

const (
	ItemStackRequestActionTypeTake                     ItemStackRequestActionType = 0
	ItemStackRequestActionTypePlace                    ItemStackRequestActionType = 1
	ItemStackRequestActionTypeSwap                     ItemStackRequestActionType = 2
	ItemStackRequestActionTypeDrop                     ItemStackRequestActionType = 3
	ItemStackRequestActionTypeDestroy                  ItemStackRequestActionType = 4
	ItemStackRequestActionTypeConsume                  ItemStackRequestActionType = 5
	ItemStackRequestActionTypeCreate                   ItemStackRequestActionType = 6
	ItemStackRequestActionTypeScreenLabTableCombine    ItemStackRequestActionType = 9
	ItemStackRequestActionTypeScreenBeaconPayment      ItemStackRequestActionType = 10
	ItemStackRequestActionTypeScreenHUDMineBlock       ItemStackRequestActionType = 11
	ItemStackRequestActionTypeCraftRecipe              ItemStackRequestActionType = 12
	ItemStackRequestActionTypeCraftRecipeAuto          ItemStackRequestActionType = 13
	ItemStackRequestActionTypeCraftCreative            ItemStackRequestActionType = 14
	ItemStackRequestActionTypeCraftRecipeOptional      ItemStackRequestActionType = 15
	ItemStackRequestActionTypeCraftRepairAndDisenchant ItemStackRequestActionType = 16
	ItemStackRequestActionTypeCraftLoom                ItemStackRequestActionType = 17
	ItemStackRequestActionTypeCraftReservedAction      ItemStackRequestActionType = 18
	ItemStackRequestActionTypeCraftNonImplemented      ItemStackRequestActionType = 19
	ItemStackRequestActionTypeCraftResults             ItemStackRequestActionType = 20
)
