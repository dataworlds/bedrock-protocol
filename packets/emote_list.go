package packets

import "github.com/dataworlds/bedrock-protocol/types"

const EmoteListID uint32 = 152

type EmoteList struct {
	RuntimeId     types.ActorRuntimeID
	EmotePieceIds []types.MceUUID
}
