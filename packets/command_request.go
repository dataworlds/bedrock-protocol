package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CommandRequestID uint32 = 77

type CommandRequest struct {
	Command    string
	Origin     types.CommandOriginData
	IsInternal bool
	Version    types.CurrentCmdVersion
}
