package packets

const SettingsCommandID uint32 = 140

type SettingsCommand struct {
	Command        string
	SuppressOutput bool
}
