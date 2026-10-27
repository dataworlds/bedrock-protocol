package packets

const GameTestResultsID uint32 = 195

type GameTestResults struct {
	Succeeded bool
	Error     string
	TestName  string
}
