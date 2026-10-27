package types

type ControlScheme uint8

const (
	ControlSchemeLockedPlayerRelativeStrafe ControlScheme = 0
	ControlSchemeCameraRelative             ControlScheme = 1
	ControlSchemeCameraRelativeStrafe       ControlScheme = 2
	ControlSchemePlayerRelative             ControlScheme = 3
	ControlSchemePlayerRelativeStrafe       ControlScheme = 4
)
