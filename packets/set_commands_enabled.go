package packets

const SetCommandsEnabledID uint32 = 59

type SetCommandsEnabled struct {
	CommandsEnabled bool
}
