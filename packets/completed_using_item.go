package packets

const CompletedUsingItemID uint32 = 142

type CompletedUsingItem struct {
	ItemId        int16
	ItemUseMethod int32
}
