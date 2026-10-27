package types

type DataDrivenScreenClosedReason uint8

const (
	DataDrivenScreenClosedReasonProgrammaticClose    DataDrivenScreenClosedReason = 0
	DataDrivenScreenClosedReasonProgrammaticCloseAll DataDrivenScreenClosedReason = 1
	DataDrivenScreenClosedReasonClientCanceled       DataDrivenScreenClosedReason = 2
	DataDrivenScreenClosedReasonUserBusy             DataDrivenScreenClosedReason = 3
	DataDrivenScreenClosedReasonInvalidForm          DataDrivenScreenClosedReason = 4
)
