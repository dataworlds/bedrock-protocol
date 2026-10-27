package types

type ActorSwingSource uint8

const (
	ActorSwingSourceNone      ActorSwingSource = 0
	ActorSwingSourceBuild     ActorSwingSource = 1
	ActorSwingSourceMine      ActorSwingSource = 2
	ActorSwingSourceInteract  ActorSwingSource = 3
	ActorSwingSourceAttack    ActorSwingSource = 4
	ActorSwingSourceUseItem   ActorSwingSource = 5
	ActorSwingSourceThrowItem ActorSwingSource = 6
	ActorSwingSourceDropItem  ActorSwingSource = 7
	ActorSwingSourceEvent     ActorSwingSource = 8
)
