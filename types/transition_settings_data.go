package types

type TransitionSettingsData struct {
	TotalTransitionTicks   uint32
	CurrentTransitionTicks uint32
	Easing                 EasingFunction
	ClockName              string
}
