package packets

import "github.com/dataworlds/bedrock-protocol/types"

const BossEventID uint32 = 74

type BossEvent struct {
	TargetActorID types.ActorUniqueID
	EventType     types.BossEventUpdateType
	Name          string
	FilteredName  string
	HealthPercent float32
	Color         types.BossBarColor
	Overlay       types.BossBarOverlay
}
