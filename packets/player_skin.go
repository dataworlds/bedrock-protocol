package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlayerSkinID uint32 = 93

type PlayerSkin struct {
	UUID                 types.MceUUID
	SerializedSkin       types.SerializedSkin
	LocalizedNewSkinName string
	LocalizedOldSkinName string
}
