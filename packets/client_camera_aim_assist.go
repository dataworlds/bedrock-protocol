package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientCameraAimAssistID uint32 = 321

type ClientCameraAimAssist struct {
	CameraPresetId string
	Action         types.ClientCameraAimAssistPacketAction
	AllowAimAssist *bool
}
