package mysql

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/morikuni/failure"

	"AmazingTalker/pkg/domain"
	"AmazingTalker/pkg/technical/errutil"
)

func NewTutorRepo(db *sqlx.DB) *TutorRepo {
	cols := []string{"t.id as tutor_id", "t.slug", "t.name", "t.headline", "t.introduction", "lang.language_id"}
	return &TutorRepo{db: db, selectCols: cols}
}

type TutorRepo struct {
	db         *sqlx.DB
	selectCols []string
	dataMapper tutorDataMapper
}

func (repo TutorRepo) QueryByTutorSlug(tutorSlug string) (domain.Tutor, error) {
	sqlString, args, _ := repo.sqlByTutorSlug(tutorSlug).ToSql()
	t := make([]*Tutor, 0)

	err := sqlx.Select(repo.db, &t, sqlString, args...)
	if err != nil {
		return domain.Tutor{}, failure.Translate(err, errutil.ErrDB)
	}

	return repo.dataMapper.toDomainTutorByTutorSlug(t), nil
}

func (repo TutorRepo) QueryAllByLanguageID(id domain.LanguageID) (domain.TutorGroup, error) {
	sqlString, args, _ := repo.sqlByLanguageID(id).ToSql()
	t := make([]*Tutor, 0)

	err := sqlx.Select(repo.db, &t, sqlString, args...)
	if err != nil {
		return nil, failure.Translate(err, errutil.ErrDB)
	}

	return repo.dataMapper.toDomainTutorGroupByLangID(t), nil
}

func (repo TutorRepo) sqlByTutorSlug(tutorSlug string) sq.SelectBuilder {
	return repo.joinSQL().Columns(repo.selectCols...).Where(sq.Eq{"t.slug": tutorSlug})
}

func (repo TutorRepo) sqlByLanguageID(id domain.LanguageID) sq.SelectBuilder {
	subQ := repo.
		joinSQL().
		Columns("t.id").
		Where(sq.Eq{"lang.language_id": id})

	return repo.
		joinSQL().
		Columns(repo.selectCols...).
		Where(SubQueryIN("t.id", subQ))
}

func (repo TutorRepo) joinSQL() sq.SelectBuilder {
	return sq.
		Select().
		From(TableNameTutor + " as t").
		InnerJoin(TableNameTutorLanguages + " as lang on t.id = lang.tutor_id")
}

func SubQueryIN(property string, query sq.SelectBuilder) sq.Sqlizer {
	sql, args, _ := query.ToSql()
	subQuery := fmt.Sprintf("%s IN (%s)", property, sql)
	return sq.Expr(subQuery, args...)
}
