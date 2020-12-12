package mysql

import (
	"github.com/jmoiron/sqlx"

	"AmazingTalker/pkg/domain"
)

func NewLanguageMapTableRepo(db *sqlx.DB) *LanguageMapTableRepo {
	return &LanguageMapTableRepo{db: db}
}

type LanguageMapTableRepo struct {
	db *sqlx.DB
}

func (repo LanguageMapTableRepo) LanguageMapTable() (domain.LanguageMapTable, error) {
	table := domain.LanguageMapTable{
		MapSlugToID: map[domain.LanguageSlug]domain.LanguageID{
			"english": 123,
		},
	}
	return table, nil
}
