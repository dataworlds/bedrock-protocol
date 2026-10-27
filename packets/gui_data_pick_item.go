package packets

const GuiDataPickItemID uint32 = 54

type GuiDataPickItem struct {
	ItemName       string
	ItemEffectName string
	Slot           int32
}
