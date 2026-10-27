package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetScoreboardIdentityID uint32 = 112

type SetScoreboardIdentity struct {
	ScoreboardIdentityPacketType uint8
	ScoreboardIdentityInfo       []types.ScoreboardIdentityPacketInfo
}
