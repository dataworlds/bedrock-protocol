package types

type CameraInstructionSplineInstruction struct {
	TotalTime         float32
	Type              uint8
	Curve             []Vec3
	ProgressKeyFrames []CameraInstructionSplineInstructionSplineProgressOption
	RotationOption    []CameraInstructionSplineInstructionSplineRotationOption
	SplineIdentifier  *string
	LoadFromJson      *bool
}
