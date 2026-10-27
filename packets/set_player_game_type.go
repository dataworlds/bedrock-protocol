package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetPlayerGameTypeID uint32 = 62

type SetPlayerGameType struct {
	PlayerGameType types.GameType
}
