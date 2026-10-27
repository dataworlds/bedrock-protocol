package packets

import "github.com/dataworlds/bedrock-protocol/types"

const UpdateSoftEnumID uint32 = 114

type UpdateSoftEnum struct {
	EnumName   string
	Values     []string
	UpdateType types.SoftEnumUpdateType
}
