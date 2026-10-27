package types

type ClientPlayMode uint32

const (
	ClientPlayModeNormal    ClientPlayMode = 0
	ClientPlayModeTeaser    ClientPlayMode = 1
	ClientPlayModeScreen    ClientPlayMode = 2
	ClientPlayModeExitLevel ClientPlayMode = 7
	ClientPlayModeNumModes  ClientPlayMode = 9
)
