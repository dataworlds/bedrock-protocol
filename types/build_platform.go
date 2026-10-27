package types

type BuildPlatform int32

const (
	BuildPlatformGoogle    BuildPlatform = 1
	BuildPlatformIOS       BuildPlatform = 2
	BuildPlatformOSX       BuildPlatform = 3
	BuildPlatformAmazon    BuildPlatform = 4
	BuildPlatformWin32     BuildPlatform = 8
	BuildPlatformDedicated BuildPlatform = 9
	BuildPlatformSony      BuildPlatform = 11
	BuildPlatformNintendo  BuildPlatform = 12
	BuildPlatformXbox      BuildPlatform = 13
	BuildPlatformLinux     BuildPlatform = 15
	BuildPlatformUnknown   BuildPlatform = -1
)
