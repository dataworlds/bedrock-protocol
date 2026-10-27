package types

type ScoreboardIdentityPacketType uint8

const (
	ScoreboardIdentityPacketTypeUpdate ScoreboardIdentityPacketType = 0
	ScoreboardIdentityPacketTypeRemove ScoreboardIdentityPacketType = 1
)
