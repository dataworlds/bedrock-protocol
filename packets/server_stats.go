package packets

const ServerStatsID uint32 = 192

type ServerStats struct {
	ServerTime  float32
	NetworkTime float32
}
