package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlayerVideoCaptureID uint32 = 324

type PlayerVideoCapture struct {
	StartVideoCapture *struct {
		ControlValue      uint32
		StartVideoCapture *types.StartVideoCapture
		StopVideoCapture  *types.StopVideoCapture
	}
	StopVideoCapture *struct {
		ControlValue      uint32
		StartVideoCapture *types.StartVideoCapture
		StopVideoCapture  *types.StopVideoCapture
	}
}
