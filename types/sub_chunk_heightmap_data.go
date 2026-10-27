package types

type SubChunkHeightmapData struct {
	HeightMapType           uint8
	SubchunkHeightMap       *[][]int8
	RenderHeightMapType     uint8
	SubchunkRenderHeightMap *[][]int8
}
