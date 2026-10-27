package packets

import "github.com/dataworlds/bedrock-protocol/types"

const TransferID uint32 = 85

type Transfer struct {
	ServerAddress           string
	ServerPort              uint16
	ReloadWorld             bool
	GatheringsConfiguration *types.GatheringsConfig
}
