package packets

import "github.com/dataworlds/bedrock-protocol/types"

const EducationSettingsID uint32 = 137

type EducationSettings struct {
	EducationLevelSettings types.EducationLevelSettings
}
