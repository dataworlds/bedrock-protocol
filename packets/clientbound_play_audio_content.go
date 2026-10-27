package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ClientboundPlayAudioContentID uint32 = 359

type ClientboundPlayAudioContent struct {
	SharedMetadata  types.SignedAudioContent
	PlaybackContent types.SignedAudioContent
	PlaybackType    types.AudioContentPlaybackType
	PlaySound       PlaySound
}
