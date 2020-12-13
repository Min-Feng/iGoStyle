// +build integration

package mysql_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"

	"AmazingTalker/pkg/domain"
	"AmazingTalker/pkg/technical/injector"
	"AmazingTalker/pkg/technical/logger"
)

func TestIntegrationLessonRepo(t *testing.T) {
	suite.Run(t, new(LessonRepoTestSuiteIntegration))
}

type LessonRepoTestSuiteIntegration struct {
	suite.Suite
}

func (ts *LessonRepoTestSuiteIntegration) Test_QueryByTutorID() {
	logger.DebugMode()
	db := injector.InfraPart()
	_, repo, _ := injector.Infra(db)

	actualLesson, err := repo.QueryByTutorID(1)

	ts.Assert().NoError(err)
	log.Debug().Msgf("\n%v", spew.Sdump(actualLesson))
}

func (ts *LessonRepoTestSuiteIntegration) Test_QueryAllByTutorIDGroup() {
	logger.DebugMode()
	db := injector.InfraPart()
	_, repo, _ := injector.Infra(db)

	actualLesson, err := repo.QueryAllByTutorIDGroup([]domain.TutorID{1, 2})

	ts.Assert().NoError(err)
	log.Debug().Msgf("\n%v", spew.Sdump(actualLesson))
}
