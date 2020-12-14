package domain

type LessonRepo interface {
	QueryByTutorID(id TutorID) (*Lesson, error)
	QueryAllByTutorIDGroup(ids []TutorID) ([]*Lesson, error)
}

type Lesson struct {
	LessonID int     `db:"id"`
	TutorID  TutorID `db:"tutor_id"`
	LessonPrice
}

type LessonPrice struct {
	Trial  float64 `json:"trial" db:"trial_price"`
	Normal float64 `json:"normal" db:"normal_price"`
}
