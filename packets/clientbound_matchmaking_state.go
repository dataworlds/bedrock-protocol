package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientboundMatchmakingStateID uint32 = 353

type ClientboundMatchmakingState struct {
	State           types.MatchmakingState
	DestinationName string
	Options         *types.MatchmakingStateOptions
}
