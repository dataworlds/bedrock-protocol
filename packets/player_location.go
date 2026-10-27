package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlayerLocationID uint32 = 326

type PlayerLocation struct {
	TargetActorID types.ActorUniqueID
	Location      struct {
		ControlValue        uint32
		CoordinatesLocation *types.CoordinatesLocation
		HiddenLocation      *types.HiddenLocation
	}
}
