package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdateTradeID uint32 = 80

type UpdateTrade struct {
	ContainerId       uint8
	Type              uint8
	Size              int32
	TraderTier        int32
	EntityUniqueId    types.ActorUniqueID
	LastTradingPlayer types.ActorUniqueID
	DisplayName       string
	UseNewTradeScreen bool
	UsingEconomyTrade bool
	Data              any
}
