package types

type AvailableCommandsPacketParamData struct {
	Name        string
	ParseSymbol uint32
	IsOptional  bool
	Options     uint8
}
