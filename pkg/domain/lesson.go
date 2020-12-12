package domain

type LessonRepo interface {
	QueryByTutorID(id TutorID) (Lesson, error)
	QueryAllByTutorIDGroup(ids []TutorID) ([]*Lesson, error)
}

type Lesson struct {
	LessonID int
	TutorID  TutorID
	LessonPrice
}

type LessonPrice struct {
	Trial  float64 `json:"trial"`
	Normal float64 `json:"normal"`
}
