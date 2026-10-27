package types

type ModalFormCancelReason uint8

const (
	ModalFormCancelReasonUserClosed ModalFormCancelReason = 0
	ModalFormCancelReasonUserBusy   ModalFormCancelReason = 1
)
