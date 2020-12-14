package mysql

import (
	"github.com/jmoiron/sqlx"

	"iGoStyle/pkg/domain"
)

func NewLanguageMapTableRepo(db *sqlx.DB) *LanguageMapTableRepo {
	return &LanguageMapTableRepo{db: db}
}

type LanguageMapTableRepo struct {
	db *sqlx.DB
}

func (repo LanguageMapTableRepo) LanguageLookupForm() (domain.LanguageLookupForm, error) {
	table := domain.LanguageLookupForm{
		MapSlugToID: map[domain.LanguageSlug]domain.LanguageID{
			"english": 123,
		},
	}
	return table, nil
}
