package types

type PacketViolationSeverity int32

const (
	PacketViolationSeverityUnknown               PacketViolationSeverity = -1
	PacketViolationSeverityWarning               PacketViolationSeverity = 0
	PacketViolationSeverityFinalWarning          PacketViolationSeverity = 1
	PacketViolationSeverityTerminatingConnection PacketViolationSeverity = 2
)
