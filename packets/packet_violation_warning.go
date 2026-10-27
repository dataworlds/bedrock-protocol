package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PacketViolationWarningID uint32 = 156

type PacketViolationWarning struct {
	ViolationType     types.PacketViolationType
	ViolationSeverity types.PacketViolationSeverity
	ViolationPacketId int32
	ViolationContext  string
}
