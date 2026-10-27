package types

type BiomeSurfaceBuilderData struct {
	SurfaceMaterials           *BiomeSurfaceMaterialData
	HasDefaultOverworldSurface bool
	HasSwampSurface            bool
	HasFrozenOceanSurface      bool
	HasTheEndSurface           bool
	MesaSurface                *BiomeMesaSurfaceData
	CappedSurface              *BiomeCappedSurfaceData
	NoiseGradientSurface       *BiomeNoiseGradientSurfaceData
}
