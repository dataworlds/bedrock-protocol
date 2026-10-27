package types

type SubChunkPacketPayloadSubChunkRequestResult uint8

const (
	SubChunkPacketPayloadSubChunkRequestResultUndefined             SubChunkPacketPayloadSubChunkRequestResult = 0
	SubChunkPacketPayloadSubChunkRequestResultSuccess               SubChunkPacketPayloadSubChunkRequestResult = 1
	SubChunkPacketPayloadSubChunkRequestResultLevelChunkDoesntExist SubChunkPacketPayloadSubChunkRequestResult = 2
	SubChunkPacketPayloadSubChunkRequestResultWrongDimension        SubChunkPacketPayloadSubChunkRequestResult = 3
	SubChunkPacketPayloadSubChunkRequestResultPlayerDoesntExist     SubChunkPacketPayloadSubChunkRequestResult = 4
	SubChunkPacketPayloadSubChunkRequestResultIndexOutOfBounds      SubChunkPacketPayloadSubChunkRequestResult = 5
	SubChunkPacketPayloadSubChunkRequestResultSuccessAllAir         SubChunkPacketPayloadSubChunkRequestResult = 6
)
