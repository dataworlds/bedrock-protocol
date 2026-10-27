package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlayerListID uint32 = 63

type PlayerList struct {
	Entries *[]struct {
		ControlValue          uint32
		PlayerListRemoveEntry *types.PlayerListRemoveEntry
		PlayerListAddEntry    *types.PlayerListAddEntry
	}
}
