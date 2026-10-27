package types

type BoolAttributeOperation int32

const (
	BoolAttributeOperationOVERRIDE   BoolAttributeOperation = 0
	BoolAttributeOperationALPHABLEND BoolAttributeOperation = 1
	BoolAttributeOperationAND        BoolAttributeOperation = 2
	BoolAttributeOperationNAND       BoolAttributeOperation = 3
	BoolAttributeOperationOR         BoolAttributeOperation = 4
	BoolAttributeOperationNOR        BoolAttributeOperation = 5
	BoolAttributeOperationXOR        BoolAttributeOperation = 6
	BoolAttributeOperationXNOR       BoolAttributeOperation = 7
)
