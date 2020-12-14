package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/morikuni/failure"

	"iGoStyle/pkg/domain"
	"iGoStyle/pkg/technical/errutil"
)

func NewLessonRepo(db *sqlx.DB) *LessonRepo {
	cols := []string{"id", "tutor_id", "trial_price", "normal_price"}
	return &LessonRepo{db: db, selectCols: cols}
}

type LessonRepo struct {
	db         *sqlx.DB
	selectCols []string
}

func (repo LessonRepo) QueryByTutorID(id domain.TutorID) (*domain.Lesson, error) {
	sqlString, args, _ := repo.sqlByTutorID(id).ToSql()

	lesson := new(domain.Lesson)
	err := sqlx.Get(repo.db, lesson, sqlString, args...)
	if err != nil {
		return nil, failure.Translate(err, errutil.ErrDB)
	}

	return lesson, nil
}

func (repo LessonRepo) sqlByTutorID(id domain.TutorID) sq.SelectBuilder {
	return sq.
		Select(repo.selectCols...).
		From(TableNameTutorLessonPrices).
		Where(sq.Eq{"tutor_id": id})
}

func (repo LessonRepo) QueryAllByTutorIDGroup(ids []domain.TutorID) ([]*domain.Lesson, error) {
	sqlString, args, _ := repo.sqlByTutorIDGroup(ids).ToSql()

	lessons := make([]*domain.Lesson, 0)
	err := sqlx.Select(repo.db, &lessons, sqlString, args...)
	if err != nil {
		return nil, failure.Translate(err, errutil.ErrDB)
	}

	return lessons, nil
}

func (repo LessonRepo) sqlByTutorIDGroup(ids []domain.TutorID) sq.SelectBuilder {
	return sq.
		Select(repo.selectCols...).
		From(TableNameTutorLessonPrices).
		Where(sq.Eq{"tutor_id": ids})
}
