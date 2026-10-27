package types

type PlayerListAddEntry struct {
	Action           uint8
	UUID             MceUUID
	ActorUniqueID    ActorUniqueID
	PlayerName       string
	XBLXUID          string
	PlayFabID        string
	PlatformOnlineID string
	BuildPlatform    BuildPlatform
	SerializedSkin   SerializedSkin
	IsTeacher        bool
	IsHost           bool
	IsSubClient      bool
	PlayerColor      Color
}
