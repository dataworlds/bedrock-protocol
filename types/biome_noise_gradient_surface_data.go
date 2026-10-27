package types

type BiomeNoiseGradientSurfaceData struct {
	NonReplaceableBlocks []uint32
	GradientBlocks       []SerializedNoiseBlockSpecifier
	Noise                NoiseDescriptor
}
