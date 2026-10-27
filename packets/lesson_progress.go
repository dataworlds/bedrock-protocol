package packets

const LessonProgressID uint32 = 183

type LessonProgress struct {
	LessonAction int32
	Score        int32
	ActivityId   string
}
