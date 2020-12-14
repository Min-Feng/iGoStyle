package mysql

import "iGoStyle/pkg/domain"

type Tutor struct {
	ID           domain.TutorID    `db:"tutor_id"`
	Slug         string            `db:"slug"`
	Name         string            `db:"name"`
	Headline     string            `db:"headline"`
	Introduction string            `db:"introduction"`
	LanguageID   domain.LanguageID `db:"language_id"`
}

type tutorDataMapper struct{}

func (m tutorDataMapper) toDomainTutorByTutorSlug(data []*Tutor) *domain.Tutor {
	tutor := new(domain.Tutor)
	for i, t := range data {
		if i == 0 {
			tutor = &domain.Tutor{
				ID:                   t.ID,
				Slug:                 t.Slug,
				Name:                 t.Name,
				Headline:             t.Headline,
				Introduction:         t.Introduction,
				ProfessionalLanguage: []domain.LanguageID{t.LanguageID},
			}
			continue
		}
		if t.ID != tutor.ID {
			panic("developer error")
		}
		tutor.ProfessionalLanguage = append(tutor.ProfessionalLanguage, t.LanguageID)
	}
	return tutor
}

func (m tutorDataMapper) toDomainTutorGroupByLangID(data []*Tutor) domain.TutorGroup {
	tutors := make(domain.TutorGroup)
	for _, t := range data {
		tutor, ok := tutors[t.ID]
		if !ok {
			tutors[t.ID] = &domain.Tutor{
				ID:                   t.ID,
				Slug:                 t.Slug,
				Name:                 t.Name,
				Headline:             t.Headline,
				Introduction:         t.Introduction,
				ProfessionalLanguage: []domain.LanguageID{t.LanguageID},
			}
		} else {
			tutor.ProfessionalLanguage = append(tutor.ProfessionalLanguage, t.LanguageID)
		}
	}
	return tutors
}
