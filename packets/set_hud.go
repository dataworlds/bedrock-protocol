package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetHudID uint32 = 308

type SetHud struct {
	HudElement []types.HudElement
	HudVisible types.HudVisibility
}
