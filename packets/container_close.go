package packets

const ContainerCloseID uint32 = 47

type ContainerClose struct {
	ContainerId          uint8
	ContainerType        uint8
	ServerInitiatedClose bool
}
