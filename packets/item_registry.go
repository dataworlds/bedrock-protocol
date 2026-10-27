package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ItemRegistryID uint32 = 162

type ItemRegistry struct {
	ItemData []types.ItemData
}
