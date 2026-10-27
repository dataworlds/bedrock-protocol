package types

type BiomeDefinitionChunkGenData struct {
	Climate                    *BiomeClimateData
	ConsolidatedFeatures       *BiomeConsolidatedFeaturesData
	MountainParams             *BiomeMountainParamsData
	SurfaceMaterialAdjustments *BiomeSurfaceMaterialAdjustmentData
	OverworldGenRules          *BiomeOverworldGenRulesData
	MultinoiseGenRules         *BiomeMultinoiseGenRulesData
	LegacyWorldGenRules        *BiomeLegacyWorldGenRulesData
	ReplacementBiomes          *BiomeReplacementsData
	VillageType                *VillageType
	SurfaceBuilderData         *BiomeSurfaceBuilderData
	SubsurfaceBuilderData      *BiomeSurfaceBuilderData
}
