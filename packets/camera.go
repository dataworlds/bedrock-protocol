package packets

import "github.com/dataworlds/bedrock-protocol/types"

const CameraID uint32 = 73

type Camera struct {
	CameraID       types.ActorUniqueID
	TargetPlayerID types.ActorUniqueID
}
