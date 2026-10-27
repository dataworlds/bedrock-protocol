package types

type ColorAttributeOperation int32

const (
	ColorAttributeOperationOVERRIDE   ColorAttributeOperation = 0
	ColorAttributeOperationALPHABLEND ColorAttributeOperation = 1
	ColorAttributeOperationADD        ColorAttributeOperation = 2
	ColorAttributeOperationSUBTRACT   ColorAttributeOperation = 3
	ColorAttributeOperationMULTIPLY   ColorAttributeOperation = 4
)
