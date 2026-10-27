package types

type DataStoreChange struct {
	DataStoreName       string
	Property            string
	UpdateCount         uint32
	TheNewPropertyValue any
}
