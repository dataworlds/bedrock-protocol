package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdateAbilitiesID uint32 = 187

type UpdateAbilities struct {
	Data types.SerializedAbilitiesData
}
