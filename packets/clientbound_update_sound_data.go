package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientboundUpdateSoundDataID uint32 = 348

type ClientboundUpdateSoundData struct {
	ServerSoundHandle types.ServerSoundHandle
	Event             struct {
		ControlValue uint32
		Stop         *types.Stop
		SetVolume    *types.SetVolume
		SetPitch     *types.SetPitch
		Fade         *types.Fade
		SeekTo       *types.SeekTo
		Pause        *types.Pause
		Resume       *types.Resume
	}
}
