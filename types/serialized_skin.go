package types

type SerializedSkin struct {
	ID                           string
	ResourcePatch                string
	ImageData                    SkinImage
	AnimatedImageData            []AnimatedImageData
	CapeImageData                SkinImage
	GeometryData                 string
	GeometryDataMinEngineVersion MinEngineVersion
	AnimationData                string
	CapeID                       string
	FullID                       string
	ArmSize                      PersonaArmSizeType
	SkinColor                    Color
	PersonaPieces                []SerializedPersonaPieceHandle
	PieceTintColors              map[PersonaPieceType]TintMapColor
	IsPremium                    bool
	IsPersona                    bool
	IsPersonaCapeOnClassicSkin   bool
	IsPrimaryUser                bool
	OverridesPlayerAppearance    bool
	TrustedSkinFlag              TrustedSkinFlag
	ProfileHash                  string
}
