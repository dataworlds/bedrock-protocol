package packets

import "github.com/dataworlds/bedrock-protocol/types"

const TextID uint32 = 9

type Text struct {
	Localize bool
	Body     struct {
		ControlValue     uint32
		MessageOnly      *types.MessageOnly
		AuthorAndMessage *types.AuthorAndMessage
		MessageAndParams *types.MessageAndParams
	}
	SenderSXUID     *string
	PlatformId      *string
	FilteredMessage *string
}
