// +build integration

package mysql_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"

	"AmazingTalker/pkg/domain"
	"AmazingTalker/pkg/infra/mysql"
	"AmazingTalker/pkg/technical/configs"
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
	cfg := configs.NewConfig()
	db, _ := injector.InfraPart(cfg)
	lessonRepo := mysql.NewLessonRepo(db)

	actualLesson, err := lessonRepo.QueryByTutorID(1)

	ts.Assert().NoError(err)
	log.Debug().Msgf("\n%v", spew.Sdump(actualLesson))
}

func (ts *LessonRepoTestSuiteIntegration) Test_QueryAllByTutorIDGroup() {
	logger.DebugMode()
	cfg := configs.NewConfig()
	db, _ := injector.InfraPart(cfg)
	lessonRepo := mysql.NewLessonRepo(db)

	actualLesson, err := lessonRepo.QueryAllByTutorIDGroup([]domain.TutorID{1, 2})

	ts.Assert().NoError(err)
	log.Debug().Msgf("\n%v", spew.Sdump(actualLesson))
}
