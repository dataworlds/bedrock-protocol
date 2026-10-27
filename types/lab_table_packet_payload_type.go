package types

type LabTablePacketPayloadType uint8

const (
	LabTablePacketPayloadTypeStartCombine  LabTablePacketPayloadType = 0
	LabTablePacketPayloadTypeStartReaction LabTablePacketPayloadType = 1
	LabTablePacketPayloadTypeReset         LabTablePacketPayloadType = 2
)
