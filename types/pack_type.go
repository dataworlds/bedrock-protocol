package types

type PackType string

const (
	PackTypeInvalid       PackType = "Invalid"
	PackTypeAddon         PackType = "Addon"
	PackTypeCached        PackType = "Cached"
	PackTypeCopyProtected PackType = "CopyProtected"
	PackTypeBehavior      PackType = "Behavior"
	PackTypePersonaPiece  PackType = "PersonaPiece"
	PackTypeResources     PackType = "Resources"
	PackTypeSkins         PackType = "Skins"
	PackTypeWorldTemplate PackType = "WorldTemplate"
)
