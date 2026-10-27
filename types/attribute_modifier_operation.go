package types

type AttributeModifierOperation int32

const (
	AttributeModifierOperationOPERATIONADDITION      AttributeModifierOperation = 0
	AttributeModifierOperationOPERATIONMULTIPLYBASE  AttributeModifierOperation = 1
	AttributeModifierOperationOPERATIONMULTIPLYTOTAL AttributeModifierOperation = 2
	AttributeModifierOperationOPERATIONCAP           AttributeModifierOperation = 3
	AttributeModifierOperationTOTALOPERATIONS        AttributeModifierOperation = 4
	AttributeModifierOperationOPERATIONINVALID       AttributeModifierOperation = 4
)
