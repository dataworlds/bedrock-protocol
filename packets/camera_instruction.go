package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CameraInstructionID uint32 = 300

type CameraInstruction struct {
	CameraInstruction types.CameraInstruction
}
