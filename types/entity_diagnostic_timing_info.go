package types

type EntityDiagnosticTimingInfo struct {
	DisplayName    string
	Entity         string
	TimeInNS       uint64
	PercentOfTotal uint8
	Position       *Vec3
	Dimension      *string
}
