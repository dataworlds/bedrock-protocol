package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetScoreID uint32 = 108

type SetScore struct {
	ScoreInfo *[]struct {
		ControlValue          uint32
		RemoveScore           *types.RemoveScore
		ChangePlayerScore     *types.ChangePlayerScore
		ChangeEntityScore     *types.ChangeEntityScore
		ChangeFakePlayerScore *types.ChangeFakePlayerScore
	}
}
