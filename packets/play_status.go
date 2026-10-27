package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlayStatusID uint32 = 2

type PlayStatus struct {
	Status types.PlayStatus
}
