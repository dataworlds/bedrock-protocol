package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CorrectPlayerMovePredictionID uint32 = 161

type CorrectPlayerMovePrediction struct {
	PredictionType         types.RewindType
	Pos                    types.Vec3
	PosDelta               types.Vec3
	Rotation               types.Vec2
	VehicleAngularVelocity *float32
	OnGround               bool
	Tick                   types.PlayerInputTick
}
