package types

type ItemUseInventoryTransaction struct {
	Actions                  InventoryTransaction
	ActionType               ItemUseActionType
	TriggerType              ItemUseTriggerType
	Position                 BlockPos
	Face                     uint8
	Slot                     int32
	Hand                     HandSlot
	Item                     NetworkItemStackDescriptor
	FromPosition             Vec3
	ClickPosition            Vec3
	TargetBlockId            uint32
	ClientInteractPrediction ItemUsePredictedResult
	ClientCooldownState      ItemUseClientCooldownState
}
