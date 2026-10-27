package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ServerboundDataStoreID uint32 = 332

type ServerboundDataStore struct {
	Update types.DataStoreUpdate
}
