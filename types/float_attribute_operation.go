package types

type FloatAttributeOperation int32

const (
	FloatAttributeOperationOVERRIDE   FloatAttributeOperation = 0
	FloatAttributeOperationALPHABLEND FloatAttributeOperation = 1
	FloatAttributeOperationADD        FloatAttributeOperation = 2
	FloatAttributeOperationSUBTRACT   FloatAttributeOperation = 3
	FloatAttributeOperationMULTIPLY   FloatAttributeOperation = 4
	FloatAttributeOperationMINIMUM    FloatAttributeOperation = 5
	FloatAttributeOperationMAXIMUM    FloatAttributeOperation = 6
)
