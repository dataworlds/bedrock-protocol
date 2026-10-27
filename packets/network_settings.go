package packets

import "github.com/dataworlds/bedrock-protocol/types"

const NetworkSettingsID uint32 = 143

type NetworkSettings struct {
	CompressionThreshold    uint16
	CompressionAlgorithm    types.PacketCompressionAlgorithm
	ClientThrottleEnabled   bool
	ClientThrottleThreshold uint8
	ClientThrottleScalar    float32
}
