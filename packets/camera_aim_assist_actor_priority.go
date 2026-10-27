package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CameraAimAssistActorPriorityID uint32 = 339

type CameraAimAssistActorPriority struct {
	CameraAimAssistActorPriorityList []types.AimAssistActorPriorityData
}
