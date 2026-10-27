package types

type ResourcePackResponse int8

const (
	ResourcePackResponseCancel                    ResourcePackResponse = 1
	ResourcePackResponseDownloading               ResourcePackResponse = 2
	ResourcePackResponseDownloadingFinished       ResourcePackResponse = 3
	ResourcePackResponseResourcePackStackFinished ResourcePackResponse = 4
)
