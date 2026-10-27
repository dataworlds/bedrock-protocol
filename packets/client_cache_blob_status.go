package packets

const ClientCacheBlobStatusID uint32 = 135

type ClientCacheBlobStatus struct {
	MissingIds []uint64
	FoundIds   []uint64
}
