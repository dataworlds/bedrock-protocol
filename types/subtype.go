package types

type Subtype uint16

const (
	SubtypeUninitializedSubtype        Subtype = 0
	SubtypeEnableCommands              Subtype = 1
	SubtypeDisableCommands             Subtype = 2
	SubtypeUnlockWorldTemplateSettings Subtype = 3
)
