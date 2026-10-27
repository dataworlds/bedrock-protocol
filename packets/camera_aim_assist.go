package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CameraAimAssistID uint32 = 316

type CameraAimAssist struct {
	PresetId        string
	ViewAngle       types.Vec2
	Distance        float32
	TargetMode      types.CameraAimAssistPacketPayloadTargetMode
	Action          types.CameraAimAssistPacketPayloadAction
	ShowDebugRender bool
}
