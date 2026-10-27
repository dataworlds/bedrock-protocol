package packets

import "github.com/dataworlds/bedrock-protocol/types"

const AnimateEntityID uint32 = 158

type AnimateEntity struct {
	MAnimation             string
	MNextState             string
	MStopExpression        string
	MStopExpressionVersion int32
	MController            string
	MBlendOutTime          float32
	MRuntimeIds            []types.ActorRuntimeID
}
