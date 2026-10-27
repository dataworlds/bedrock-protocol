package types

type ItemData struct {
	ItemName          HashedString
	ItemId            int16
	IsComponentBased  bool
	ItemVersion       ItemVersion
	ItemComponentData any
}
