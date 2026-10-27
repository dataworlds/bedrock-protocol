package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlayerEnchantOptionsID uint32 = 146

type PlayerEnchantOptions struct {
	Options []types.ItemEnchantOption
}
