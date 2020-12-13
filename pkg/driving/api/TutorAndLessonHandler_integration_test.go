// +build integration

package api_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"

	"AmazingTalker/pkg/technical/configs"
	"AmazingTalker/pkg/technical/injector"
	"AmazingTalker/pkg/technical/logger"
	"AmazingTalker/pkg/technical/testutil"
)

func TestTutorAndLessonHandler(t *testing.T) {
	suite.Run(t, new(TutorAndLessonHandlerTestSuite))
}

type TutorAndLessonHandlerTestSuite struct {
	suite.Suite
}

func (ts *TutorAndLessonHandlerTestSuite) SetupTest() {
}

func (ts *TutorAndLessonHandlerTestSuite) Test_GetTutors() {
	logger.DebugMode()
	cfg := configs.NewConfig()
	router := injector.Project(cfg)

	languageSlug := "english"
	urlString := "/api/tutors/" + languageSlug

	var actualBody string
	for i := 0; i < 5; i++ {
		actualBody = testutil.HTTPResponseBody(router, http.MethodGet, urlString, nil)
	}
	log.Debug().Msgf("\n%v", actualBody)
	_ = actualBody
}

func (ts *TutorAndLessonHandlerTestSuite) Test_GetTutor() {
	logger.DebugMode()
	cfg := configs.NewConfig()
	router := injector.Project(cfg)

	tutorSlug, _ := url.Parse("at-1")
	urlString := "/api/tutor/" + tutorSlug.String()

	var actualBody string
	for i := 0; i < 5; i++ {
		actualBody = testutil.HTTPResponseBody(router, http.MethodGet, urlString, nil)
	}
	log.Debug().Msgf("\n%v", actualBody)
	_ = actualBody
}
