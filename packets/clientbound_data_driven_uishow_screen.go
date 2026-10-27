package packets

const ClientboundDataDrivenUIShowScreenID uint32 = 333

type ClientboundDataDrivenUIShowScreen struct {
	ScreenId       string
	FormId         uint32
	DataInstanceId *uint32
}
