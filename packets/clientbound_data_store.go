package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientboundDataStoreID uint32 = 330

type ClientboundDataStore struct {
	Updates []struct {
		ControlValue     uint32
		DataStoreUpdate  *types.DataStoreUpdate
		DataStoreChange  *types.DataStoreChange
		DataStoreRemoval *types.DataStoreRemoval
	}
}
