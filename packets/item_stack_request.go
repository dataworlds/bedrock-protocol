package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ItemStackRequestID uint32 = 147

type ItemStackRequest struct {
	Requests []types.ItemStackRequest
}
