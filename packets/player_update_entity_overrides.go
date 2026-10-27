package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlayerUpdateEntityOverridesID uint32 = 325

type PlayerUpdateEntityOverrides struct {
	TargetID      types.ActorUniqueID
	PropertyIndex uint32
	Update        struct {
		ControlValue   uint32
		ClearOverride  *types.ClearOverride
		RemoveOverride *types.RemoveOverride
		IntOverride    *types.IntOverride
		FloatOverride  *types.FloatOverride
	}
}
