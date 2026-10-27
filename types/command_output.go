package types

type CommandOutput struct {
	OutputType     CommandOutputType
	SuccessCount   uint32
	OutputMessages *[]CommandOutputMessage
	DataSet        *string
}
