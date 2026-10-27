package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CommandOutputID uint32 = 79

type CommandOutput struct {
	OriginData types.CommandOriginData
	Output     types.CommandOutput
}
