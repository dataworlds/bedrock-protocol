package packets

const LoginID uint32 = 1

type Login struct {
	ClientNetworkVersion int32
	ConnectionRequest    string
}
