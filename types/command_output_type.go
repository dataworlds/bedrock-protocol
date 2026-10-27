package types

type CommandOutputType int32

const (
	CommandOutputTypeNone       CommandOutputType = 0
	CommandOutputTypeLastOutput CommandOutputType = 1
	CommandOutputTypeSilent     CommandOutputType = 2
	CommandOutputTypeAllOutput  CommandOutputType = 3
	CommandOutputTypeDataSet    CommandOutputType = 4
)
