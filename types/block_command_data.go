package types

type BlockCommandData struct {
	BlockPosition    BlockPos
	CommandBlockMode uint32
	RedstoneMode     bool
	IsConditional    bool
}
