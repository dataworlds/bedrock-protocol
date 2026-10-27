package packets

const ServerboundLoadingScreenID uint32 = 312

type ServerboundLoadingScreen struct {
	LoadingScreenPacketType int32
	LoadingScreenId         *uint32
}
