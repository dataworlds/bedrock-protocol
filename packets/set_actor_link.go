package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetActorLinkID uint32 = 41

type SetActorLink struct {
	Link types.ActorLink
}
