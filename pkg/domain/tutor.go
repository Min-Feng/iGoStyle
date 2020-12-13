package domain

import "sort"

type TutorRepo interface {
	QueryByTutorSlug(tutorSlug string) (*Tutor, error)
	QueryAllByLanguageID(id LanguageID) (TutorGroup, error)
}

type Tutor struct {
	ID                   TutorID      `json:"id,string"`
	Slug                 string       `json:"slug"`
	Name                 string       `json:"name"`
	Headline             string       `json:"headline"`
	Introduction         string       `json:"introduction"`
	ProfessionalLanguage []LanguageID `json:"teaching_languages"`
}

type TutorID int
type TutorGroup map[TutorID]*Tutor

func (tutors TutorGroup) IDGroup() []TutorID {
	ids := make([]TutorID, 0, len(tutors))
	for _, tutor := range tutors {
		ids = append(ids, tutor.ID)
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	return ids
}
