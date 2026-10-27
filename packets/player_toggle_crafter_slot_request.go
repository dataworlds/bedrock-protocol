package packets

const PlayerToggleCrafterSlotRequestID uint32 = 306

type PlayerToggleCrafterSlotRequest struct {
	PosX       int32
	PosY       int32
	PosZ       int32
	SlotIndex  uint8
	IsDisabled bool
}
