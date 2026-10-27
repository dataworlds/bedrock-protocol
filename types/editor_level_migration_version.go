package types

type EditorLevelMigrationVersion int8

const (
	EditorLevelMigrationVersionLegacy             EditorLevelMigrationVersion = 0
	EditorLevelMigrationVersionEditorAchievements EditorLevelMigrationVersion = 1
)
