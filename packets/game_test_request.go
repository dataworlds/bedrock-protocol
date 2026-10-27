package packets

import "github.com/dataworlds/bedrock-protocol/types"

const GameTestRequestID uint32 = 194

type GameTestRequest struct {
	MaxTestsPerBatch int32
	RepeatCount      int32
	Rotation         types.Rotation
	StopOnFailure    bool
	TestPos          types.BlockPos
	TestsPerRow      int32
	TestName         string
}
