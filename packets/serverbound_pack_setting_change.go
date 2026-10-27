package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ServerboundPackSettingChangeID uint32 = 329

type ServerboundPackSettingChange struct {
	PackId           types.MceUUID
	PackSettingName  string
	PackSettingValue struct {
		ControlValue uint32
		Variant0     *float32
		Variant1     *bool
		Variant2     *string
		Variant3     *[]string
	}
}
