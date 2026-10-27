package types

type CommandOutputMessage struct {
	MessageID  string
	Successful bool
	Parameters *[]string
}
