package application

import (
	"encoding/json"

	"AmazingTalker/pkg/domain"
)

type TutorAndLessonViewFactory struct{}

func (f TutorAndLessonViewFactory) createArrayData(lessons []*domain.Lesson, tutors domain.TutorGroup) []byte {
	views := make([]TutorAndLessonViewModel, 0, len(lessons))
	for _, lesson := range lessons {
		tutor := tutors[lesson.TutorID]
		view := f.createViewModel(lesson, tutor)
		views = append(views, view)
	}
	data, _ := json.Marshal(&TutorAndLessonArrayData{views})
	return data
}

type TutorAndLessonArrayData struct {
	Data []TutorAndLessonViewModel `json:"data"`
}

func (f TutorAndLessonViewFactory) createSingleData(lesson *domain.Lesson, tutor *domain.Tutor) []byte {
	view := f.createViewModel(lesson, tutor)
	data, _ := json.Marshal(&TutorAndLessonSingleData{view})
	return data
}

type TutorAndLessonSingleData struct {
	TutorAndLessonViewModel `json:"data"`
}

func (f TutorAndLessonViewFactory) createViewModel(lesson *domain.Lesson, tutor *domain.Tutor) TutorAndLessonViewModel {
	return TutorAndLessonViewModel{
		Tutor: *tutor,
		Price: domain.LessonPrice{
			Trial:  lesson.LessonPrice.Trial,
			Normal: lesson.LessonPrice.Normal,
		},
	}
}

type TutorAndLessonViewModel struct {
	domain.Tutor
	Price domain.LessonPrice `json:"price_info"`
}
