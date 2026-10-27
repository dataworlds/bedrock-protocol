package types

type ChangePlayerScore struct {
	Action         ScorePacketEntryAction
	ScoreboardId   ScoreboardId
	ObjectiveName  string
	ScoreValue     int32
	PlayerUniqueId PlayerScoreboardId
}
