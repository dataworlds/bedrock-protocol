package types

type ServerConfig struct {
	Gathering             *GatheringsConfig
	ClientStoreEntryPoint *ClientStoreEntryPointConfig
	Presence              *PresenceConfig
}
