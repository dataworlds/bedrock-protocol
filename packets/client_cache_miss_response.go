package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientCacheMissResponseID uint32 = 136

type ClientCacheMissResponse struct {
	MissingBlobs []types.MissingBlobData
}
