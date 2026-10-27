package types

type CameraPresets struct {
	Name                           SharedTypesIdentifierSharedTypesAssetTypeCameraPreset
	InheritFrom                    SharedTypesReferenceSharedTypesAssetTypeCameraPreset
	PosX                           *float32
	PosY                           *float32
	PosZ                           *float32
	RotX                           *float32
	RotY                           *float32
	RotationSpeed                  *float32
	SnapToTarget                   *bool
	HorizontalRotationLimit        *Vec2
	VerticalRotationLimit          *Vec2
	ContinueTargeting              *bool
	BlockListeningRadius           *float32
	ViewOffset                     *Vec2
	EntityOffset                   *Vec3
	Radius                         *float32
	YawLimitMin                    *float32
	YawLimitMax                    *float32
	Listener                       *SharedTypesComprehensiveCameraPresetAudioListener
	PlayerEffects                  *bool
	AimAssist                      *SharedTypesComprehensiveCameraAimAssistCommandDefinition
	ControlScheme                  *ControlScheme
	ApplyInheritedStartingRotation bool
	StartingRotation               *Vec2
}
