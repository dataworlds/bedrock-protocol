package packets

import "github.com/dataworlds/bedrock-protocol/types"

const MotionPredictionHintsID uint32 = 157

type MotionPredictionHints struct {
	MRuntimeId types.ActorRuntimeID
	MMotion    types.Vec3
	MOnGround  bool
}
