package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CameraAimAssistPresetsID uint32 = 320

type CameraAimAssistPresets struct {
	CameraAimAssistPresets    []types.CameraAimAssistCategoryDefinition
	CameraAimAssistCategories []types.CameraAimAssistPresetDefinition
	Operation                 types.CameraAimAssistPresetsPacketOperation
}
