package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ResourcePackClientResponseID uint32 = 8

type ResourcePackClientResponse struct {
	Response struct {
		ControlValue                                        uint32
		ResourcePackClientResponseCancel                    *types.ResourcePackClientResponseCancel
		ResourcePackClientResponseDownloading               *types.ResourcePackClientResponseDownloading
		ResourcePackClientResponseDownloadingFinished       *types.ResourcePackClientResponseDownloadingFinished
		ResourcePackClientResponseResourcePackStackFinished *types.ResourcePackClientResponseResourcePackStackFinished
	}
}
