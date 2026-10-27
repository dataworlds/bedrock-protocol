package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ServerPresenceInfoID uint32 = 347

type ServerPresenceInfo struct {
	PresenceConfiguration *types.PresenceConfig
}
