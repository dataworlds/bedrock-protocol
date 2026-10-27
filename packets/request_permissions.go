package packets

const RequestPermissionsID uint32 = 185

type RequestPermissions struct {
	TargetPlayerIdSRawID  int64
	PlayerPermissionLevel int32
	CustomPermissionFlags uint16
}
