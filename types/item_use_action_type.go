package types

type ItemUseActionType int32

const (
	ItemUseActionTypePlace       ItemUseActionType = 0
	ItemUseActionTypeUse         ItemUseActionType = 1
	ItemUseActionTypeDestroy     ItemUseActionType = 2
	ItemUseActionTypeUseAsAttack ItemUseActionType = 3
)
