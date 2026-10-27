package types

type ChatRestrictionLevel uint8

const (
	ChatRestrictionLevelNone     ChatRestrictionLevel = 0
	ChatRestrictionLevelDropped  ChatRestrictionLevel = 1
	ChatRestrictionLevelDisabled ChatRestrictionLevel = 2
)
