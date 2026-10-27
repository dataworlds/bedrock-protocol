package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientMovementPredictionSyncID uint32 = 322

type ClientMovementPredictionSync struct {
	ActorDataFlag      types.ActorDataFlagComponent
	ActorBoundingBox   types.ActorDataBoundingBoxComponent
	MovementAttributes types.MovementAttributesComponent
	ActorUniqueID      types.ActorUniqueID
	ActorFlyingState   bool
}
