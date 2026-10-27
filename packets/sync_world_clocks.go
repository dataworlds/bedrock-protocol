package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SyncWorldClocksID uint32 = 344

type SyncWorldClocks struct {
	Data struct {
		ControlValue           uint32
		SyncStateData          *types.SyncStateData
		InitializeRegistryData *types.InitializeRegistryData
		AddTimeMarkerData      *types.AddTimeMarkerData
		RemoveTimeMarkerData   *types.RemoveTimeMarkerData
	}
}
