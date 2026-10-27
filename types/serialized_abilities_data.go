package types

type SerializedAbilitiesData struct {
	TargetPlayerRawId  int64
	PlayerPermissions  PlayerPermissionLevel
	CommandPermissions CommandPermissionLevel
	Layers             []SerializedAbilitiesDataSerializedLayer
}
