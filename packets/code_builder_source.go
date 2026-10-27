package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CodeBuilderSourceID uint32 = 178

type CodeBuilderSource struct {
	Operation  types.CodeBuilderStorageQueryOptionsOperation
	Category   types.CodeBuilderStorageQueryOptionsCategory
	CodeStatus types.CodeBuilderExecutionStateCodeStatus
}
