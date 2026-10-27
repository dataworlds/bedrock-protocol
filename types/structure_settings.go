package types

type StructureSettings struct {
	StructurePaletteName                            string
	ShouldIgnoreEntities                            bool
	ShouldIgnoreBlocks                              bool
	ShouldAllowNonTickingPlayerAndTickingAreaChunks bool
	StructureSize                                   BlockPos
	StructureOffset                                 BlockPos
	LastEditPlayer                                  ActorUniqueID
	Rotation                                        Rotation
	Mirror                                          Mirror
	AnimationMode                                   AnimationMode
	AnimationSeconds                                float32
	IntegrityValue                                  float32
	IntegritySeed                                   uint32
	RotationPivot                                   Vec3
}
