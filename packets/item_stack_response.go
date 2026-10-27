package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ItemStackResponseID uint32 = 148

type ItemStackResponse struct {
	Responses []types.ItemStackResponseInfo
}
