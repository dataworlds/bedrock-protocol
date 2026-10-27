package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetPlayerInventoryOptionsID uint32 = 307

type SetPlayerInventoryOptions struct {
	InventoryOptions types.InventoryOptions
}
