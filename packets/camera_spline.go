package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CameraSplineID uint32 = 338

type CameraSpline struct {
	CameraDataSplines []types.CameraSplineDefinition
}
