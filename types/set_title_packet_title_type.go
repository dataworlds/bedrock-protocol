package types

type SetTitlePacketTitleType int32

const (
	SetTitlePacketTitleTypeClear               SetTitlePacketTitleType = 0
	SetTitlePacketTitleTypeReset               SetTitlePacketTitleType = 1
	SetTitlePacketTitleTypeTitle               SetTitlePacketTitleType = 2
	SetTitlePacketTitleTypeSubtitle            SetTitlePacketTitleType = 3
	SetTitlePacketTitleTypeActionbar           SetTitlePacketTitleType = 4
	SetTitlePacketTitleTypeTimes               SetTitlePacketTitleType = 5
	SetTitlePacketTitleTypeTitleTextObject     SetTitlePacketTitleType = 6
	SetTitlePacketTitleTypeSubtitleTextObject  SetTitlePacketTitleType = 7
	SetTitlePacketTitleTypeActionbarTextObject SetTitlePacketTitleType = 8
)
