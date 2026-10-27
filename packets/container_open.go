package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ContainerOpenID uint32 = 46

type ContainerOpen struct {
	ContainerId   uint8
	ContainerType uint8
	Position      types.BlockPos
	TargetActorID types.ActorUniqueID
}
