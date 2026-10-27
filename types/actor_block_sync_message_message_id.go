package types

type ActorBlockSyncMessageMessageId uint32

const (
	ActorBlockSyncMessageMessageIdNONE    ActorBlockSyncMessageMessageId = 0
	ActorBlockSyncMessageMessageIdCREATE  ActorBlockSyncMessageMessageId = 1
	ActorBlockSyncMessageMessageIdDESTROY ActorBlockSyncMessageMessageId = 2
)
