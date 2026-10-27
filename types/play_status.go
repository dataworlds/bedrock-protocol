package types

type PlayStatus int32

const (
	PlayStatusLoginSuccess                             PlayStatus = 0
	PlayStatusLoginFailedClientOld                     PlayStatus = 1
	PlayStatusLoginFailedServerOld                     PlayStatus = 2
	PlayStatusPlayerSpawn                              PlayStatus = 3
	PlayStatusLoginFailedInvalidTenant                 PlayStatus = 4
	PlayStatusLoginFailedEditionMismatchEduToVanilla   PlayStatus = 5
	PlayStatusLoginFailedEditionMismatchVanillaToEdu   PlayStatus = 6
	PlayStatusLoginFailedServerFullSubClient           PlayStatus = 7
	PlayStatusLoginFailedEditorMismatchEditorToVanilla PlayStatus = 8
	PlayStatusLoginFailedEditorMismatchVanillaToEditor PlayStatus = 9
)
