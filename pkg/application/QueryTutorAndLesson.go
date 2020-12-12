package application

import (
	"github.com/morikuni/failure"

	"AmazingTalker/pkg/domain"
)

func NewTutorAndLessonUseCase(
	lessonRepo domain.LessonRepo,
	tutorRepo domain.TutorRepo,
	langRep domain.LanguageMapTableRepo,
) *TutorAndLessonUseCase {

	languageMapTable, err := langRep.LanguageMapTable()
	if err != nil {
		panic(err)
	}

	return &TutorAndLessonUseCase{
		lessonRepo:   lessonRepo,
		tutorRepo:    tutorRepo,
		langMapTable: languageMapTable,
	}
}

type TutorAndLessonUseCase struct {
	lessonRepo   domain.LessonRepo
	tutorRepo    domain.TutorRepo
	langMapTable domain.LanguageMapTable
	viewFactory  TutorAndLessonViewFactory
}

func (uc *TutorAndLessonUseCase) QueryByLanguageSlug(langSlug string) ([]byte, error) {
	languageID := uc.langMapTable.SlugToID(domain.LanguageSlug(langSlug))

	tutors, err := uc.tutorRepo.QueryAllByLanguageID(languageID)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	lessons, err := uc.lessonRepo.QueryAllByTutorIDGroup(tutors.IDGroup())
	if err != nil {
		return nil, failure.Wrap(err)
	}

	return uc.viewFactory.createArrayData(lessons, tutors), nil
}

func (uc *TutorAndLessonUseCase) QueryByTutorSlug(tutorSlug string) ([]byte, error) {
	tutor, err := uc.tutorRepo.QueryByTutorSlug(tutorSlug)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	lesson, err := uc.lessonRepo.QueryByTutorID(tutor.ID)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	return uc.viewFactory.createSingleData(&lesson, &tutor), nil
}
