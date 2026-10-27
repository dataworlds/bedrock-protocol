package types

type SoundDataEventType uint8

const (
	SoundDataEventTypeStop      SoundDataEventType = 0
	SoundDataEventTypeSetVolume SoundDataEventType = 1
	SoundDataEventTypeSetPitch  SoundDataEventType = 2
	SoundDataEventTypeFade      SoundDataEventType = 3
	SoundDataEventTypeSeekTo    SoundDataEventType = 4
	SoundDataEventTypePause     SoundDataEventType = 5
	SoundDataEventTypeResume    SoundDataEventType = 6
)
