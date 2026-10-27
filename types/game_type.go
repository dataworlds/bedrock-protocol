package types

type GameType int32

const (
	GameTypeUndefined    GameType = -1
	GameTypeSurvival     GameType = 0
	GameTypeCreative     GameType = 1
	GameTypeAdventure    GameType = 2
	GameTypeDefault      GameType = 5
	GameTypeSpectator    GameType = 6
	GameTypeWorldDefault GameType = 0
)
