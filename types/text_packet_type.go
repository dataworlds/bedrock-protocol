package types

type TextPacketType uint8

const (
	TextPacketTypeRaw                    TextPacketType = 0
	TextPacketTypeChat                   TextPacketType = 1
	TextPacketTypeTranslate              TextPacketType = 2
	TextPacketTypePopup                  TextPacketType = 3
	TextPacketTypeJukeboxPopup           TextPacketType = 4
	TextPacketTypeTip                    TextPacketType = 5
	TextPacketTypeSystemMessage          TextPacketType = 6
	TextPacketTypeWhisper                TextPacketType = 7
	TextPacketTypeAnnouncement           TextPacketType = 8
	TextPacketTypeTextObjectWhisper      TextPacketType = 9
	TextPacketTypeTextObject             TextPacketType = 10
	TextPacketTypeTextObjectAnnouncement TextPacketType = 11
)
