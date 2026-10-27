package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PhotoTransferID uint32 = 99

type PhotoTransfer struct {
	PhotoName    string
	PhotoData    string
	BookID       string
	Type         types.PhotoType
	SourceType   types.PhotoType
	OwnerID      int64
	NewPhotoName string
}
