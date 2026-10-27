package types

type ServerboundLoadingScreenPacketType int32

const (
	ServerboundLoadingScreenPacketTypeUnknown            ServerboundLoadingScreenPacketType = 0
	ServerboundLoadingScreenPacketTypeStartLoadingScreen ServerboundLoadingScreenPacketType = 1
	ServerboundLoadingScreenPacketTypeEndLoadingScreen   ServerboundLoadingScreenPacketType = 2
)
