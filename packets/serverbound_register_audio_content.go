package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ServerboundRegisterAudioContentID uint32 = 360

type ServerboundRegisterAudioContent struct {
	Registrations []types.AudioContentRegistrationEntry
}
