package types

type CommandBlockMode uint16

const (
	CommandBlockModeNormal    CommandBlockMode = 0
	CommandBlockModeRepeating CommandBlockMode = 1
	CommandBlockModeChain     CommandBlockMode = 2
)
