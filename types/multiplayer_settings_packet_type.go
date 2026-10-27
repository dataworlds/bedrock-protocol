package types

type MultiplayerSettingsPacketType int32

const (
	MultiplayerSettingsPacketTypeEnableMultiplayer  MultiplayerSettingsPacketType = 0
	MultiplayerSettingsPacketTypeDisableMultiplayer MultiplayerSettingsPacketType = 1
	MultiplayerSettingsPacketTypeRefreshJoincode    MultiplayerSettingsPacketType = 2
)
