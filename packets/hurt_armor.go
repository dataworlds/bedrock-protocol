package packets

const HurtArmorID uint32 = 38

type HurtArmor struct {
	Cause      int32
	Damage     int32
	ArmorSlots uint64
}
