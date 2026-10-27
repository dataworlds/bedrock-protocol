package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetTitleID uint32 = 88

type SetTitle struct {
	TitleType            types.SetTitlePacketTitleType
	TitleText            string
	FadeInTime           int32
	StayTime             int32
	FadeOutTime          int32
	Xuid                 string
	PlatformOnlineId     string
	FilteredTitleMessage string
}
