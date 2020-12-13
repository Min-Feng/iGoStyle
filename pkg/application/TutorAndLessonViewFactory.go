package application

import (
	"encoding/json"

	"AmazingTalker/pkg/domain"
)

type tutorAndLessonViewFactory struct{}

func (f tutorAndLessonViewFactory) createArrayData(lessons []*domain.Lesson, tutors domain.TutorGroup) []byte {
	views := make([]tutorAndLessonViewModel, 0, len(lessons))
	for _, lesson := range lessons {
		tutor := tutors[lesson.TutorID]
		view := f.createViewModel(lesson, tutor)
		views = append(views, view)
	}
	data, _ := json.Marshal(&tutorAndLessonArrayData{views})
	return data
}

type tutorAndLessonArrayData struct {
	Data []tutorAndLessonViewModel `json:"data"`
}

func (f tutorAndLessonViewFactory) createSingleData(lesson *domain.Lesson, tutor *domain.Tutor) []byte {
	view := f.createViewModel(lesson, tutor)
	data, _ := json.Marshal(&tutorAndLessonSingleData{view})
	return data
}

type tutorAndLessonSingleData struct {
	tutorAndLessonViewModel `json:"data"`
}

func (f tutorAndLessonViewFactory) createViewModel(lesson *domain.Lesson, tutor *domain.Tutor) tutorAndLessonViewModel {
	return tutorAndLessonViewModel{
		Tutor: *tutor,
		Price: domain.LessonPrice{
			Trial:  lesson.LessonPrice.Trial,
			Normal: lesson.LessonPrice.Normal,
		},
	}
}

type tutorAndLessonViewModel struct {
	domain.Tutor
	Price domain.LessonPrice `json:"price_info"`
}
