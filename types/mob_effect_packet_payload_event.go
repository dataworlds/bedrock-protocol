package types

type MobEffectPacketPayloadEvent uint8

const (
	MobEffectPacketPayloadEventInvalid MobEffectPacketPayloadEvent = 0
	MobEffectPacketPayloadEventAdd     MobEffectPacketPayloadEvent = 1
	MobEffectPacketPayloadEventUpdate  MobEffectPacketPayloadEvent = 2
	MobEffectPacketPayloadEventRemove  MobEffectPacketPayloadEvent = 3
)
