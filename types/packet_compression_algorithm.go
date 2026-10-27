package types

type PacketCompressionAlgorithm uint16

const (
	PacketCompressionAlgorithmZLib   PacketCompressionAlgorithm = 0
	PacketCompressionAlgorithmSnappy PacketCompressionAlgorithm = 1
	PacketCompressionAlgorithmNone   PacketCompressionAlgorithm = 65535
)
