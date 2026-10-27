package packets

import "github.com/dataworlds/bedrock-protocol/types"

const SetPassengerOfBlockID uint32 = 357

type SetPassengerOfBlock struct {
	Passenger            types.ActorUniqueID
	PassengerOfBlockData *types.PassengerOfBlockArguments
}
