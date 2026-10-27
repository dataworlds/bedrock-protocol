package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ServerToClientHandshakeID uint32 = 3

type ServerToClientHandshake struct {
	HandshakeWebToken types.WebToken
}
