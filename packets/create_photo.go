package packets

const CreatePhotoID uint32 = 171

type CreatePhoto struct {
	RawID         uint64
	PhotoName     string
	PhotoItemName string
}
