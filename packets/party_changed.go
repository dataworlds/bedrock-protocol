package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PartyChangedID uint32 = 342

type PartyChanged struct {
	PartyInfo *types.PlayerPartyInfo
}
