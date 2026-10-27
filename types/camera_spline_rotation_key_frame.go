package types

type CameraSplineRotationKeyFrame struct {
	Rotation Vec3
	Time     float32
	Easing   *EasingFunction
}
