package types

type PackInfoData struct {
	PackIdVersion       PackIdVersion
	PackSize            uint64
	ContentKey          string
	SubpackName         string
	ContentIdentity     ContentIdentity
	HasScripts          bool
	IsAddonPack         bool
	IsRayTracingCapable bool
	CDNURL              string
}
