package packets

import "github.com/dataworlds/bedrock-protocol/types"

const GameRulesChangedID uint32 = 72

type GameRulesChanged struct {
	RuleData types.GameRulesChangedPacketData
}
