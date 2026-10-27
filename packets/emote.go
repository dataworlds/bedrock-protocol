package packets

import "github.com/dataworlds/bedrock-protocol/types"

const EmoteID uint32 = 138

type Emote struct {
	ActorRuntimeId   types.ActorRuntimeID
	EmoteId          string
	EmoteLengthTicks uint32
	Xuid             string
	PlatformId       string
	Flags            uint8
}
