package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ShowCreditsID uint32 = 75

type ShowCredits struct {
	PlayerRuntimeID types.ActorRuntimeID
	CreditsState    int32
}
