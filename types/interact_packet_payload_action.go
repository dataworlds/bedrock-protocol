package types

type InteractPacketPayloadAction uint8

const (
	InteractPacketPayloadActionInvalid        InteractPacketPayloadAction = 0
	InteractPacketPayloadActionStopRiding     InteractPacketPayloadAction = 3
	InteractPacketPayloadActionInteractUpdate InteractPacketPayloadAction = 4
	InteractPacketPayloadActionNpcOpen        InteractPacketPayloadAction = 5
	InteractPacketPayloadActionOpenInventory  InteractPacketPayloadAction = 6
)
