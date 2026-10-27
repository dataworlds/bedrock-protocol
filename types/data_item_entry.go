package types

type DataItemEntry struct {
	ID      uint32
	Payload struct {
		ControlValue               uint32
		DataItemBytePayload        *DataItemBytePayload
		DataItemShortPayload       *DataItemShortPayload
		DataItemIntPayload         *DataItemIntPayload
		DataItemFloatPayload       *DataItemFloatPayload
		DataItemStringPayload      *DataItemStringPayload
		DataItemCompoundTagPayload *DataItemCompoundTagPayload
		DataItemPosPayload         *DataItemPosPayload
		DataItemInt64Payload       *DataItemInt64Payload
		DataItemVec3Payload        *DataItemVec3Payload
	}
}
