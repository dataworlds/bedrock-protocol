package packets

import "github.com/dataworlds/bedrock-protocol/types"

const RequestAbilityID uint32 = 184

type RequestAbility struct {
	Ability   int32
	ValueType types.RequestAbilityPacketPayloadType
	Bool      bool
	Float     float32
}
