package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CommandBlockUpdateID uint32 = 78

type CommandBlockUpdate struct {
	Target struct {
		ControlValue        uint32
		EntityCommandTarget *types.EntityCommandTarget
		BlockCommandData    *types.BlockCommandData
	}
	Command            string
	LastOutput         string
	Name               string
	FilteredName       string
	TrackOutput        bool
	TickDelay          int32
	ExecuteOnFirstTick bool
}
