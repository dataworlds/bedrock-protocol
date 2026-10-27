package types

type ItemStackResponseSlotInfo struct {
	RequestedSlot        uint8
	Slot                 uint8
	Amount               uint8
	ItemStackNetId       *ItemStackNetId
	CustomName           BedrockSafetyRedactableString
	DurabilityCorrection int32
}
