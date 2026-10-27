package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdateEquipID uint32 = 81

type UpdateEquip struct {
	ContainerId    uint8
	Type           uint8
	Size           int32
	EntityUniqueId types.ActorUniqueID
	Data           any
}
