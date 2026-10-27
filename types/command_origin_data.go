package types

type CommandOriginData struct {
	Type      uint8
	UUID      MceUUID
	RequestId *string
	PlayerId  *int64
}
