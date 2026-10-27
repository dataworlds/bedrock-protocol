package packets

import "github.com/dataworlds/bedrock-protocol/types"

const DeathInfoID uint32 = 189

type DeathInfo struct {
	DeathCauseMessage types.DeathCauseMessageType
}
