package packets

import "github.com/dataworlds/bedrock-protocol/types"

const PlayerAuthInputID uint32 = 144

type PlayerAuthInput struct {
	PlayerRotation         types.Vec2
	Position               types.Vec3
	MoveVector             types.Vec2
	PlayerHeadRotation     float32
	InputData              []types.PlayerAuthInputData
	InputMode              types.InputMode
	PlayMode               types.ClientPlayMode
	NewInteractionModel    types.NewInteractionModel
	InteractRotation       types.Vec2
	ClientTick             types.PlayerInputTick
	PosDelta               types.Vec3
	ItemUseTransaction     *types.PackedItemUseLegacyInventoryTransaction
	ItemStackRequest       *types.ItemStackRequest
	PlayerBlockActions     *[]types.PlayerBlockActionData
	VehicleRotation        *types.Vec2
	ClientPredictedVehicle *types.ActorUniqueID
	AnalogMoveVector       types.Vec2
	CameraOrientation      types.Vec3
	RawMoveVector          types.Vec2
}
