package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AnvilDamageID uint32 = 141

type AnvilDamage struct {
	BlockPosition types.BlockPos
}
