package types

type MessageAndParams struct {
	MessageType   uint8
	Message       *string
	ParameterList *[]string
}
