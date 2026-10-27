package types

type PacketViolationType int32

const (
	PacketViolationTypeUnknown         PacketViolationType = -1
	PacketViolationTypePacketMalformed PacketViolationType = 0
)
