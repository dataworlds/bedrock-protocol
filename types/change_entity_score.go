package types

type ChangeEntityScore struct {
	Action        ScorePacketEntryAction
	ScoreboardId  ScoreboardId
	ObjectiveName string
	ScoreValue    int32
	ActorId       ActorUniqueID
}
