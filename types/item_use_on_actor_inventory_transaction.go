package types

type ItemUseOnActorInventoryTransaction struct {
	Actions      InventoryTransaction
	RuntimeId    ActorRuntimeID
	ActionType   ItemUseOnActorActionType
	Slot         int32
	Hand         HandSlot
	Item         NetworkItemStackDescriptor
	FromPosition Vec3
	HitPosition  Vec3
}
