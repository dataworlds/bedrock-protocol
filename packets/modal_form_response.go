package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ModalFormResponseID uint32 = 101

type ModalFormResponse struct {
	FormID           uint32
	JSONResponse     *string
	FormCancelReason *types.ModalFormCancelReason
}
