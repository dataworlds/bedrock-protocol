package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CameraShakeID uint32 = 159

type CameraShake struct {
	Intensity   float32
	Seconds     float32
	ShakeType   types.CameraShakeType
	ShakeAction types.CameraShakeAction
}
