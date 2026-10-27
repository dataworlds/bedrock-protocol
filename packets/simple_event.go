package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SimpleEventID uint32 = 64

type SimpleEvent struct {
	Type types.Subtype
}
