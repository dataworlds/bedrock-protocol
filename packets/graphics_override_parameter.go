package packets

import "github.com/dataworlds/bedrock-protocol/types"

const GraphicsOverrideParameterID uint32 = 331

type GraphicsOverrideParameter struct {
	ParameterKeyframeValues map[float32]types.Vec3
	FloatValue              *float32
	Vec3Value               *types.Vec3
	BiomeIdentifier         string
	PlayerIdentifier        *string
	IdentifierForParameter  types.GraphicsOverrideParameterType
	ResetParameter          bool
}
