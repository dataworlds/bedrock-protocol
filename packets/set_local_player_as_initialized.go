package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetLocalPlayerAsInitializedID uint32 = 113

type SetLocalPlayerAsInitialized struct {
	PlayerID types.ActorRuntimeID
}
