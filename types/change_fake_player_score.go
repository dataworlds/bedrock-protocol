package types

type ChangeFakePlayerScore struct {
	Action         ScorePacketEntryAction
	ScoreboardId   ScoreboardId
	ObjectiveName  string
	ScoreValue     int32
	FakePlayerName string
}
