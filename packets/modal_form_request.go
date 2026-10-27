package packets

const ModalFormRequestID uint32 = 100

type ModalFormRequest struct {
	FormID     uint32
	FormUIJSON string
}
