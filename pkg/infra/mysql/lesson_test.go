package mysql

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"

	"AmazingTalker/pkg/domain"
	"AmazingTalker/pkg/technical/types"
)

func TestLessonRepo(t *testing.T) {
	suite.Run(t, new(LessonRepoTestSuite))
}

type LessonRepoTestSuite struct {
	suite.Suite
}

func (ts *LessonRepoTestSuite) Test_sqlByTutorID() {
	repo := NewLessonRepo(nil)
	expectedSQL := "SELECT id, tutor_id, trial_price, normal_price FROM tutor_lesson_prices WHERE tutor_id = '1'"

	sqlBuilder := repo.sqlByTutorID(1)
	actualSQL := sq.DebugSqlizer(sqlBuilder)

	ts.Assert().Equal(expectedSQL, actualSQL)
}

func (ts *LessonRepoTestSuite) Test_sqlByTutorIDGroup() {
	repo := NewLessonRepo(nil)
	expectedSQL := types.StringTool{}.ToRawSQL(`
SELECT id, tutor_id, trial_price, normal_price 
FROM tutor_lesson_prices 
WHERE tutor_id IN ('1','2','4')
`)

	sqlBuilder := repo.sqlByTutorIDGroup([]domain.TutorID{1, 2, 4})
	actualSQL := sq.DebugSqlizer(sqlBuilder)

	ts.Assert().Equal(expectedSQL, actualSQL)
}
