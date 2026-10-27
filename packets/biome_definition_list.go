package packets

import "github.com/dataworlds/bedrock-protocol/types"

const BiomeDefinitionListID uint32 = 122

type BiomeDefinitionList struct {
	MapOfBiomeNamesToData map[uint16]types.BiomeDefinitionData
	StringList            types.BiomeStringList
}
