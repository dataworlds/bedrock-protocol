package types

type EditorWorldType int32

const (
	EditorWorldTypeNonEditor          EditorWorldType = 0
	EditorWorldTypeEditorProject      EditorWorldType = 1
	EditorWorldTypeEditorTestLevel    EditorWorldType = 2
	EditorWorldTypeEditorRealmsUpload EditorWorldType = 3
)
