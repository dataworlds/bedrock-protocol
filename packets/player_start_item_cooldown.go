package packets

const PlayerStartItemCooldownID uint32 = 176

type PlayerStartItemCooldown struct {
	ItemCategory  string
	DurationTicks int32
}
