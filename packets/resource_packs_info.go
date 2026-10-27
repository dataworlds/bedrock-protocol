package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ResourcePacksInfoID uint32 = 6

type ResourcePacksInfo struct {
	ResourcePackRequired       bool
	HasAddonPacks              bool
	HasScripts                 bool
	ForceDisableVibrantVisuals bool
	WorldTemplateIdAndVersion  types.PackIdVersion
	ResourcePacks              []types.PackInfoData
}
