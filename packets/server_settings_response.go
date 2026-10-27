package packets

const ServerSettingsResponseID uint32 = 103

type ServerSettingsResponse struct {
	FormID     uint32
	FormUIJSON string
}
