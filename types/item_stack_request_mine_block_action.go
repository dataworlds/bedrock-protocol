package types

type ItemStackRequestMineBlockAction struct {
	ActionType          ItemStackRequestActionType
	Slot                int32
	PredictedDurability int32
	NetIdVariant        ItemStackNetIdVariant
}
