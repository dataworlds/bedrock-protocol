package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientboundAttributeLayerSyncID uint32 = 345

type ClientboundAttributeLayerSync struct {
	Data struct {
		ControlValue                     uint32
		UpdateAttributeLayersData        *types.UpdateAttributeLayersData
		UpdateAttributeLayerSettingsData *types.UpdateAttributeLayerSettingsData
		UpdateEnvironmentAttributesData  *types.UpdateEnvironmentAttributesData
		RemoveEnvironmentAttributesData  *types.RemoveEnvironmentAttributesData
	}
}
