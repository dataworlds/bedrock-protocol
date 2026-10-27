package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetActorDataID uint32 = 39

type SetActorData struct {
	TargetRuntimeID   types.ActorRuntimeID
	ActorData         types.SynchedActorDataList
	SynchedProperties types.PropertySyncData
	Tick              types.PlayerInputTick
}
