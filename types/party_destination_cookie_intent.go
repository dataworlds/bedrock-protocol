package types

type PartyDestinationCookieIntent int8

const (
	PartyDestinationCookieIntentNotify PartyDestinationCookieIntent = 0
	PartyDestinationCookieIntentOptIn  PartyDestinationCookieIntent = 1
	PartyDestinationCookieIntentOptOut PartyDestinationCookieIntent = 2
)
