package types

type EducationLevelSettings struct {
	CodeBuilderDefaultURI        string
	CodeBuilderTitle             string
	CanResizeCodeBuilder         bool
	DisableLegacyTitleBar        bool
	PostProcessFilter            string
	ScreenshotBorderResourcePath string
	AgentCapabilities            *AgentCapabilities
	LocalSettings                EducationLocalLevelSettings
	DeprecatedAlwaysFalse        bool
	ExternalLinkSettings         *ExternalLinkSettings
}
