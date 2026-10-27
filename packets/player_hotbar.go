package packets

const PlayerHotbarID uint32 = 48

type PlayerHotbar struct {
	SelectedSlot     uint32
	ContainerID      uint8
	ShouldSelectSlot bool
}
