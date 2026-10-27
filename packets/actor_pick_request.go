package packets

const ActorPickRequestID uint32 = 35

type ActorPickRequest struct {
	ActorID  int64
	MaxSlots uint8
	WithData bool
}
