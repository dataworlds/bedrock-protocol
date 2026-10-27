package types

type CameraInstructionSetInstruction struct {
	Preset                              uint32
	Ease                                *CameraInstructionSetInstructionEaseOption
	Pos                                 *CameraInstructionSetInstructionPosOption
	Rot                                 *CameraInstructionSetInstructionRotOption
	Facing                              *CameraInstructionSetInstructionFacingOption
	ViewOffset                          *CameraInstructionSetInstructionViewOffsetOption
	EntityOffset                        *CameraInstructionSetInstructionEntityOffsetOption
	Default                             *bool
	RemoveIgnoreStartingValuesComponent bool
}
