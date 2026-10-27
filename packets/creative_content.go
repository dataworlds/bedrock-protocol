package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CreativeContentID uint32 = 145

type CreativeContent struct {
	Groups  []types.CreativeGroupInfoPayload
	Entries []types.CreativeItemEntryPayload
}
