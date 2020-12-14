// +build integration

package api_test

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"

	"iGoStyle/pkg/technical/configs"
	"iGoStyle/pkg/technical/injector"
	"iGoStyle/pkg/technical/logger"
	"iGoStyle/pkg/technical/testutil"
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
	router := injector.Server(cfg)

	languageSlug := "english"
	urlString := "/api/tutors/" + languageSlug

	var actualBody string
	for i := 0; i < 5; i++ {
		go func() {
			actualBody = testutil.HTTPResponse(router, http.MethodGet, urlString, nil)
		}()
	}
	time.Sleep(time.Second)
	log.Debug().Msgf("\n%v", actualBody)
	_ = actualBody
}

func (ts *TutorAndLessonHandlerTestSuite) Test_GetTutor() {
	logger.DebugMode()
	cfg := configs.NewConfig()
	router := injector.Server(cfg)

	tutorSlug, _ := url.Parse("at-1")
	urlString := "/api/tutor/" + tutorSlug.String()

	var actualBody string
	for i := 0; i < 5; i++ {
		actualBody = testutil.HTTPResponse(router, http.MethodGet, urlString, nil)
	}
	log.Debug().Msgf("\n%v", actualBody)
	_ = actualBody
}
