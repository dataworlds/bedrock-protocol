package packets

const StopSoundID uint32 = 87

type StopSound struct {
	SoundName       string
	StopAllSounds   bool
	StopMusicLegacy bool
}
