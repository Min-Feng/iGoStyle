package domain

type TutorRepo interface {
	QueryByTutorSlug(tutorSlug string) (Tutor, error)
	QueryAllByLanguageID(id LanguageID) (TutorGroup, error)
}

type Tutor struct {
	ID                   TutorID      `json:"id,string" db:"tutor_id"`
	Slug                 string       `json:"slug" db:"slug"`
	Name                 string       `json:"name" db:"name"`
	Headline             string       `json:"headline" db:"headline"`
	Introduction         string       `json:"introduction" db:"introduction"`
	ProfessionalLanguage []LanguageID `json:"teaching_languages" db:"language_id"`
}

type TutorID int
type TutorGroup map[TutorID]*Tutor

func (tutors TutorGroup) IDGroup() []TutorID {
	ids := make([]TutorID, 0, len(tutors))
	for _, tutor := range tutors {
		ids = append(ids, tutor.ID)
	}
	return ids
}
