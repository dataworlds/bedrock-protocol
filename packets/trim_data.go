package packets

import "github.com/dataworlds/bedrock-protocol/types"

const TrimDataID uint32 = 302

type TrimData struct {
	TrimPatternList  []types.TrimPattern
	TrimMaterialList []types.TrimMaterial
}
