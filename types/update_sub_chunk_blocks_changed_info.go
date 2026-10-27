package types

type UpdateSubChunkBlocksChangedInfo struct {
	BlocksChangedStandards []UpdateSubChunkNetworkBlockInfo
	BlocksChangedExtras    []UpdateSubChunkNetworkBlockInfo
}
