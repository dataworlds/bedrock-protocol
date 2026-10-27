package types

type TextDataPayload struct {
	Text             string
	UseRotation      *bool
	BackgroundColor  *Color
	LineGapHeight    *float32
	DepthTest        *bool
	ShowBackface     *bool
	ShowTextBackface *bool
}
