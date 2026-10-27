package types

type ItemStackResponseInfo struct {
	Result          uint8
	ClientRequestId ItemStackRequestId
	Containers      *[]ItemStackResponseContainerInfo
}
