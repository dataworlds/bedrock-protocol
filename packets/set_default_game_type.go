package packets

const SetDefaultGameTypeID uint32 = 105

type SetDefaultGameType struct {
	DefaultGameType int32
}
