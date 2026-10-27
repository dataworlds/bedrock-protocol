package packets

const LevelEventGenericID uint32 = 124

type LevelEventGeneric struct {
	EventId int32
	CTD     any
}
