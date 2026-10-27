package types

type SerializedPersonaPieceHandle struct {
	PieceId        string
	PieceType      PersonaPieceType
	PackId         MceUUID
	IsDefaultPiece bool
	ProductId      string
}
