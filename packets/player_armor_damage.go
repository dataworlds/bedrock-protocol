package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlayerArmorDamageID uint32 = 149

type PlayerArmorDamage struct {
	ArmorSlotAndDamagePairs []types.ArmorSlotAndDamagePair
}
