package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MultiplayerSettingsID uint32 = 139

type MultiplayerSettings struct {
	PacketType types.MultiplayerSettingsPacketType
}
