package mysql

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"

	"iGoStyle/pkg/technical/types"
)

func TestTutorRepo(t *testing.T) {
	suite.Run(t, new(TutorRepoTestSuite))
}

type TutorRepoTestSuite struct {
	suite.Suite
}

func (ts *TutorRepoTestSuite) Test_sqlByTutorSlug() {
	tutorRepo := NewTutorRepo(nil)
	expectedSQL := types.StringUtil{}.ToRawSQL(`
SELECT t.id as tutor_id, t.slug, t.name, t.headline, t.introduction, lang.language_id 
FROM tutors as t 
INNER JOIN tutor_languages as lang on t.id = lang.tutor_id 
WHERE t.slug = 'at-1'
`)

	builder := tutorRepo.sqlByTutorSlug("at-1")
	actualSQL := squirrel.DebugSqlizer(builder)

	ts.Assert().Equal(expectedSQL, actualSQL)
}

func (ts *TutorRepoTestSuite) Test_sqlByLanguageID() {
	tutorRepo := NewTutorRepo(nil)
	expectedSQL := types.StringUtil{}.ToRawSQL(`
SELECT 
	t.id as tutor_id, 
	t.slug, 
	t.name, 
	t.headline, 
	t.introduction, 
	lang.language_id 
FROM 
	tutors as t 
INNER JOIN tutor_languages as lang on 
	t.id = lang.tutor_id 
WHERE 
	t.id IN (
	SELECT 
		t.id 
	FROM 
		tutors as t 
	INNER JOIN tutor_languages as lang on 
		t.id = lang.tutor_id 
	WHERE 
		lang.language_id = '123')
`)

	builder := tutorRepo.sqlByLanguageID(123)
	actualSQL := squirrel.DebugSqlizer(builder)

	ts.Assert().Equal(expectedSQL, actualSQL)
}
