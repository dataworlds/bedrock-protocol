package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ResourcePackStackID uint32 = 7

type ResourcePackStack struct {
	TexturePackRequired bool
	TexturePackList     []types.PackInstanceId
	BaseGameVersion     string
	Experiments         types.Experiments
	IncludeEditorPacks  bool
}
