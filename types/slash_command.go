package types

type SlashCommand struct {
	SuccessCount int32
	ErrorCount   int32
	CommandName  string
	ErrorList    string
}
