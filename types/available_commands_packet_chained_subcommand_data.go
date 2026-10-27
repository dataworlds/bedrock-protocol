package types

type AvailableCommandsPacketChainedSubcommandData struct {
	Name             string
	SubCommandValues *[]AvailableCommandsPacketChainedSubcommandRelationship
}
