package types

type ClientCameraAimAssistPacketAction uint8

const (
	ClientCameraAimAssistPacketActionSetFromCameraPreset ClientCameraAimAssistPacketAction = 0
	ClientCameraAimAssistPacketActionClear               ClientCameraAimAssistPacketAction = 1
)
