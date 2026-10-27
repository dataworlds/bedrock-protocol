package types

type AudioContentRegistrationEntry struct {
	AudioContentID  string
	SharedMetadata  SignedAudioContent
	ServerContent   SignedAudioContent
	PlaybackContent SignedAudioContent
}
