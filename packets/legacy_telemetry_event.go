package packets

import "github.com/dataworlds/bedrock-protocol/types"

const LegacyTelemetryEventID uint32 = 65

type LegacyTelemetryEvent struct {
	TargetActorID types.ActorUniqueID
	EventType     types.LegacyTelemetryEventPacketPayloadType
	UsePlayerID   bool
	EventData     struct {
		ControlValue               uint32
		Achievement                *types.Achievement
		Interaction                *types.Interaction
		PortalCreated              *types.PortalCreated
		PortalUsed                 *types.PortalUsed
		MobKilled                  *types.MobKilled
		CauldronUsed               *types.CauldronUsed
		PlayerDied                 *types.PlayerDied
		BossKilled                 *types.BossKilled
		SlashCommand               *types.SlashCommand
		MobBorn                    *types.MobBorn
		POICauldronUsed            *types.POICauldronUsed
		ComposterUsed              *types.ComposterUsed
		BellUsed                   *types.BellUsed
		ActorDefinition            *types.ActorDefinition
		RaidUpdate                 *types.RaidUpdate
		TargetBlockHit             *types.TargetBlockHit
		PiglinBarter               *types.PiglinBarter
		PlayerWaxedOrUnwaxedCopper *types.PlayerWaxedOrUnwaxedCopper
		CodeBuilderRuntimeAction   *types.CodeBuilderRuntimeAction
		CodeBuilderScoreboard      *types.CodeBuilderScoreboard
		ItemUsed                   *types.ItemUsed
		Empty                      *types.Empty
	}
}
