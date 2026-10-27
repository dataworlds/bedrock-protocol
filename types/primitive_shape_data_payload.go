package types

type PrimitiveShapeDataPayload struct {
	NetworkId             uint64
	ShapeType             *ScriptModuleMinecraftScriptPrimitiveShapeType
	Location              *Vec3
	Scale                 *float32
	Rotation              *Vec3
	TotalTimeLeft         *float32
	MaximumRenderDistance *float32
	Color                 *Color
	DimensionID           *DimensionType
	AttachedToEntityID    *ActorUniqueID
	ExtraShapeData        *struct {
		ControlValue         uint32
		Variant0             *any
		ArrowDataPayload     *ArrowDataPayload
		TextDataPayload      *TextDataPayload
		BoxDataPayload       *BoxDataPayload
		LineDataPayload      *LineDataPayload
		SphereDataPayload    *SphereDataPayload
		CylinderDataPayload  *CylinderDataPayload
		PyramidDataPayload   *PyramidDataPayload
		EllipsoidDataPayload *EllipsoidDataPayload
		ConeDataPayload      *ConeDataPayload
	}
}
