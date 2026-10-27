package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ServerStoreInfoID uint32 = 346

type ServerStoreInfo struct {
	ClientStoreEntryPointConfiguration *types.ClientStoreEntryPointConfig
}
