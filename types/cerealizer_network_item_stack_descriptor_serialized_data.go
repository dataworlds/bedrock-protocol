package types

type CerealizerNetworkItemStackDescriptorSerializedData struct {
	Id             int16
	StackSize      uint16
	AuxValue       uint32
	NetIdVariant   *ItemStackNetIdVariant
	BlockRuntimeId uint32
	UserDataBuffer string
}
