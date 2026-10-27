package types

type MatchmakingState uint8

const (
	MatchmakingStateIdle             MatchmakingState = 0
	MatchmakingStateMatchmaking      MatchmakingState = 1
	MatchmakingStateMatchFound       MatchmakingState = 2
	MatchmakingStateCanceled         MatchmakingState = 3
	MatchmakingStatePlayerLeftParty  MatchmakingState = 4
	MatchmakingStatePlayerLeftServer MatchmakingState = 5
	MatchmakingStateServerShutdown   MatchmakingState = 6
	MatchmakingStateTimedOut         MatchmakingState = 7
	MatchmakingStateRequeueAsParty   MatchmakingState = 8
)
