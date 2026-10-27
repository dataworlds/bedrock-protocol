package types

type NoiseTransitionSettingsData struct {
	TotalTransitionTicks   uint32
	CurrentTransitionTicks uint32
	Easing                 EasingFunction
	ClockName              string
	LocalTransitionTicks   uint32
	NoiseName              string
	NoiseAlignment         NoiseAlignment
}
