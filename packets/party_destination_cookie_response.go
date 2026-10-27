package packets

const PartyDestinationCookieResponseID uint32 = 350

type PartyDestinationCookieResponse struct {
	Cookie   string
	Accepted bool
}
