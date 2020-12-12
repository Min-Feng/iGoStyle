package mysql

import (
	"github.com/jmoiron/sqlx"

	"AmazingTalker/pkg/domain"
)

func NewTutorRepo(db *sqlx.DB) *TutorRepo {
	return &TutorRepo{db: db}
}

type TutorRepo struct {
	db *sqlx.DB
}

func (repo TutorRepo) QueryByTutorSlug(tutorSlug string) (domain.Tutor, error) {
	panic("implement me")
}

func (repo TutorRepo) QueryAllByLanguageID(id domain.LanguageID) (domain.TutorGroup, error) {
	panic("implement me")
}
