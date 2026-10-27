package types

type GatheringsConfig struct {
	ExperienceId   *MceUUID
	ExperienceName *string
	WorldId        *MceUUID
	WorldName      *string
	CreatorId      *string
	TargetId       *MceUUID
	ScenarioId     *string
	ServerId       *string
}
