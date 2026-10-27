package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AvailableCommandsID uint32 = 76

type AvailableCommands struct {
	EnumValues              []string
	ChainedSubcommandValues []string
	PostFixes               []string
	EnumData                []types.AvailableCommandsPacketEnumData
	ChainedSubcommandData   []types.AvailableCommandsPacketChainedSubcommandData
	Commands                []types.AvailableCommandsPacketCommandData
	SoftEnums               []types.AvailableCommandsPacketSoftEnumData
	Constraints             []types.AvailableCommandsPacketConstrainedValueData
}
