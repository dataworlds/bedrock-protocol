package types

type UpdateSubChunkNetworkBlockInfo struct {
	Pos                       BlockPos
	RuntimeId                 uint32
	UpdateFlags               uint32
	SyncMessageEntityUniqueID uint64
	SyncMessageMessage        uint32
}
