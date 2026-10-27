package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ShowStoreOfferID uint32 = 91

type ShowStoreOffer struct {
	OfferId      types.MceUUID
	RedirectType types.ShowStoreOfferRedirectType
}
