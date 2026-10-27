package types

type ItemUseTriggerType uint8

const (
	ItemUseTriggerTypeUnknown        ItemUseTriggerType = 0
	ItemUseTriggerTypePlayerInput    ItemUseTriggerType = 1
	ItemUseTriggerTypeSimulationTick ItemUseTriggerType = 2
)
