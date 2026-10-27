package types

type GamePublishSetting int32

const (
	GamePublishSettingNoMultiPlay      GamePublishSetting = 0
	GamePublishSettingInviteOnly       GamePublishSetting = 1
	GamePublishSettingFriendsOnly      GamePublishSetting = 2
	GamePublishSettingFriendsOfFriends GamePublishSetting = 3
	GamePublishSettingPublic           GamePublishSetting = 4
)
