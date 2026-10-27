package packets

const SetHealthID uint32 = 42

type SetHealth struct {
	Health int32
}
