package types

type BiomeCappedSurfaceData struct {
	FloorBlocks     []uint32
	CeilingBlocks   []uint32
	SeaBlock        *uint32
	FoundationBlock *uint32
	BeachBlock      *uint32
}
