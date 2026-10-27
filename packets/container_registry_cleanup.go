package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ContainerRegistryCleanupID uint32 = 317

type ContainerRegistryCleanup struct {
	RemovedContainers []types.FullContainerName
}
