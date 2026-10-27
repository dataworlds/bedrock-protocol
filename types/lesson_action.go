package types

type LessonAction int8

const (
	LessonActionStart    LessonAction = 0
	LessonActionComplete LessonAction = 1
	LessonActionRestart  LessonAction = 2
)
