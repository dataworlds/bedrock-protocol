package types

type ScorePacketEntryAction uint8

const (
	ScorePacketEntryActionRemove           ScorePacketEntryAction = 0
	ScorePacketEntryActionChangePlayer     ScorePacketEntryAction = 1
	ScorePacketEntryActionChangeEntity     ScorePacketEntryAction = 2
	ScorePacketEntryActionChangeFakePlayer ScorePacketEntryAction = 3
)
