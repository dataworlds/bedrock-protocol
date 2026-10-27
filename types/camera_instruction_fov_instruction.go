package types

type CameraInstructionFovInstruction struct {
	FieldOfView      float32
	FOVEaseTime      float32
	FOVEaseType      EasingFunction
	FieldOfViewClear bool
}
