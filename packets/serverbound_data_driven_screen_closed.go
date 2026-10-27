package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ServerboundDataDrivenScreenClosedID uint32 = 343

type ServerboundDataDrivenScreenClosed struct {
	FormId      *uint32
	CloseReason types.DataDrivenScreenClosedReason
}
