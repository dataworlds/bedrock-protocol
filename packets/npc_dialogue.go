package packets

import "github.com/dataworlds/bedrock-protocol/types"

const NpcDialogueID uint32 = 169

type NpcDialogue struct {
	NpcIdRawId            uint64
	NpcDialogueActionType types.NpcDialogueActionType
	Dialogue              string
	SceneName             string
	NpcName               string
	ActionJSON            string
}
