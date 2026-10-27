package types

type CameraSplineDefinition struct {
	Name              string
	TotalTime         float32
	SplineType        *string
	ControlPoints     []CameraSplineControlPoint
	ProgressKeyFrames []CameraSplineProgressKeyFrame
	RotationKeyFrames []CameraSplineRotationKeyFrame
}
