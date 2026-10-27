package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdateAdventureSettingsID uint32 = 188

type UpdateAdventureSettings struct {
	AdventureSettings types.AdventureSettings
}
