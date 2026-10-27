package packets

import "github.com/dataworlds/bedrock-protocol/types"

const StartGameID uint32 = 11

type StartGame struct {
	EntityID                          types.ActorUniqueID
	RuntimeID                         types.ActorRuntimeID
	GameType                          types.GameType
	Position                          types.Vec3
	Rotation                          types.Vec2
	Settings                          types.LevelSettings
	LevelID                           string
	LevelName                         string
	TemplateContentIdentity           string
	IsTrial                           bool
	MovementSettings                  types.SyncedPlayerMovementSettings
	LevelCurrentTime                  uint64
	EnchantmentSeed                   int32
	BlockProperties                   []types.ServerBlockProperty
	MultiplayerCorrelationId          string
	EnableItemStackNetManager         bool
	ServerVersion                     string
	PlayerPropertyData                any
	ServerBlockTypeRegistryChecksum   uint64
	WorldTemplateID                   types.MceUUID
	ServerEnabledClientSideGeneration bool
	BlockNetworkIdsAreHashes          bool
	NetworkPermissions                types.NetworkPermissions
	ServerConfigurationJoinInfo       *types.ServerConfig
	ServerTelemetryData               types.ServerTelemetryData
}
