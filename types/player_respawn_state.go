package types

type PlayerRespawnState uint8

const (
	PlayerRespawnStateSearchingForSpawn  PlayerRespawnState = 0
	PlayerRespawnStateReadyToSpawn       PlayerRespawnState = 1
	PlayerRespawnStateClientReadyToSpawn PlayerRespawnState = 2
)
