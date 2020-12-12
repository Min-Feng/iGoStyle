package domain

type LessonRepo interface {
	QueryByTutorID(id TutorID) (Lesson, error)
	QueryAllByTutorIDGroup(ids []TutorID) ([]*Lesson, error)
}

// 看範例 sql 以及 json
// 同一個老師, 教不同語言, 價錢是一樣的?
// 好像有點怪
type Lesson struct {
	LessonID int     `db:"id"`
	TutorID  TutorID `db:"tutor_id"`
	LessonPrice
}

type LessonPrice struct {
	Trial  float64 `json:"trial" db:"trial_price"`
	Normal float64 `json:"normal" db:"normal_price"`
}
