package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdateClientOptionsID uint32 = 323

type UpdateClientOptions struct {
	GraphicsModeChange    *types.GraphicsMode
	FilterProfanityChange *bool
}
