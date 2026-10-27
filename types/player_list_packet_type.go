package types

type PlayerListPacketType uint8

const (
	PlayerListPacketTypeAdd    PlayerListPacketType = 0
	PlayerListPacketTypeRemove PlayerListPacketType = 1
)
