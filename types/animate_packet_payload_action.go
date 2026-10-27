package types

type AnimatePacketPayloadAction uint8

const (
	AnimatePacketPayloadActionNoAction         AnimatePacketPayloadAction = 0
	AnimatePacketPayloadActionSwing            AnimatePacketPayloadAction = 1
	AnimatePacketPayloadActionWakeUp           AnimatePacketPayloadAction = 3
	AnimatePacketPayloadActionCriticalHit      AnimatePacketPayloadAction = 4
	AnimatePacketPayloadActionMagicCriticalHit AnimatePacketPayloadAction = 5
)
