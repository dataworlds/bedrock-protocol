package types

type ServerEditorConnectionPolicy int32

const (
	ServerEditorConnectionPolicyMatchWorldType ServerEditorConnectionPolicy = 0
	ServerEditorConnectionPolicyEditorOnly     ServerEditorConnectionPolicy = 1
	ServerEditorConnectionPolicyVanillaOnly    ServerEditorConnectionPolicy = 2
	ServerEditorConnectionPolicyMixed          ServerEditorConnectionPolicy = 3
)
