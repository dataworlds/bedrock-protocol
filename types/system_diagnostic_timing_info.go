package types

type SystemDiagnosticTimingInfo struct {
	DisplayName    string
	SystemIndex    uint64
	TimeInNS       uint64
	PercentOfTotal uint8
}
