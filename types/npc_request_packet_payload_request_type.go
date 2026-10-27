package types

type NpcRequestPacketPayloadRequestType uint8

const (
	NpcRequestPacketPayloadRequestTypeSetActions             NpcRequestPacketPayloadRequestType = 0
	NpcRequestPacketPayloadRequestTypeExecuteAction          NpcRequestPacketPayloadRequestType = 1
	NpcRequestPacketPayloadRequestTypeExecuteClosingCommands NpcRequestPacketPayloadRequestType = 2
	NpcRequestPacketPayloadRequestTypeSetName                NpcRequestPacketPayloadRequestType = 3
	NpcRequestPacketPayloadRequestTypeSetSkin                NpcRequestPacketPayloadRequestType = 4
	NpcRequestPacketPayloadRequestTypeSetInteractText        NpcRequestPacketPayloadRequestType = 5
	NpcRequestPacketPayloadRequestTypeExecuteOpeningCommands NpcRequestPacketPayloadRequestType = 6
)
