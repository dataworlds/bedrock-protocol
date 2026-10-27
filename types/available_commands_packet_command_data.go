package types

type AvailableCommandsPacketCommandData struct {
	Name                                string
	Description                         string
	Flags                               uint16
	PermissionLevel                     CommandPermissionLevel
	AliasEnum                           int32
	CommandDataChainedSubcommandIndexes []uint32
	Overloads                           []AvailableCommandsPacketOverloadData
}
