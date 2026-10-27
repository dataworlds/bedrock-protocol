package types

type CameraInstruction struct {
	Set              *CameraInstructionSetInstruction
	Clear            *bool
	Fade             *CameraInstructionFadeInstruction
	Target           *CameraInstructionTargetInstruction
	RemoveTarget     *bool
	FieldOfView      *CameraInstructionFovInstruction
	Spline           *CameraInstructionSplineInstruction
	AttachToEntity   *CameraInstructionAttachToEntityInstruction
	DetachFromEntity *bool
}
