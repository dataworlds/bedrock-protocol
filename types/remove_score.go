package types

type RemoveScore struct {
	Action        ScorePacketEntryAction
	ScoreboardId  ScoreboardId
	ObjectiveName *string
}
