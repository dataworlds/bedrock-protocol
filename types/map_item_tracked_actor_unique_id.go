package types

type MapItemTrackedActorUniqueId struct {
	Type          MapItemTrackedActorType
	EntityID      *ActorUniqueID
	BlockPosition *BlockPos
}
