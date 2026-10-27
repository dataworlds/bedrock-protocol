package types

type ItemStackRequestNetworkItemInstanceDescriptor struct {
	ItemDescriptor struct {
		ControlValue         uint32
		EmptyItemDescriptor  *EmptyItemDescriptor
		ItemNameDescriptor   *ItemNameDescriptor
		MolangItemDescriptor *MolangItemDescriptor
		ItemTagDescriptor    *ItemTagDescriptor
	}
	StackSize      uint16
	BlockRuntimeId uint32
	UserDataBuffer string
}
