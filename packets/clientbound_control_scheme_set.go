package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientboundControlSchemeSetID uint32 = 327

type ClientboundControlSchemeSet struct {
	ControlScheme types.ControlScheme
}
