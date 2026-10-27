package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetLastHurtByID uint32 = 96

type SetLastHurtBy struct {
	LastHurtBy types.ActorType
}
