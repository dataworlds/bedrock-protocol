package types

type Difficulty int32

const (
	DifficultyPeaceful Difficulty = 0
	DifficultyEasy     Difficulty = 1
	DifficultyNormal   Difficulty = 2
	DifficultyHard     Difficulty = 3
	DifficultyCount    Difficulty = 4
	DifficultyUnknown  Difficulty = 5
)
