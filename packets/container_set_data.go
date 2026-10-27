package packets

const ContainerSetDataID uint32 = 51

type ContainerSetData struct {
	ContainerID uint8
	ID          int32
	Value       int32
}
