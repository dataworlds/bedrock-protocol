package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ChangeMobPropertyID uint32 = 182

type ChangeMobProperty struct {
	ActorId              types.ActorUniqueID
	PropertyName         string
	BoolComponentValue   bool
	StringComponentValue string
	IntComponentValue    int32
	FloatComponentValue  float32
}
