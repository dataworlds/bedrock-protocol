package packets

const SetDisplayObjectiveID uint32 = 107

type SetDisplayObjective struct {
	DisplaySlotName      string
	ObjectiveName        string
	ObjectiveDisplayName string
	CriteriaName         string
	SortOrder            int32
}
