package types

type CreativeItemEntryPayload struct {
	CreativeNetId CreativeItemNetId
	ItemInstance  NetworkItemInstanceDescriptor
	GroupIndex    uint32
}
