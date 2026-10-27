package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SendPartyDestinationCookieID uint32 = 349

type SendPartyDestinationCookie struct {
	Cookie          string
	Intent          types.PartyDestinationCookieIntent
	DestinationName string
}
