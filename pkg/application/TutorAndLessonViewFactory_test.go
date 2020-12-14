package application

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"iGoStyle/pkg/domain"
	"iGoStyle/pkg/technical/types"
)

func TestTutorAndLessonViewFactory(t *testing.T) {
	suite.Run(t, new(TutorAndLessonViewFactoryTestSuite))
}

type TutorAndLessonViewFactoryTestSuite struct {
	suite.Suite
}

func (ts *TutorAndLessonViewFactoryTestSuite) Test_createArrayData() {
	lessons := []*domain.Lesson{
		{
			LessonID: 2,
			TutorID:  1,
			LessonPrice: domain.LessonPrice{
				Trial:  5,
				Normal: 10,
			},
		},
	}

	tutors := domain.TutorGroup{
		1: {
			ID:                   1,
			Slug:                 "at-1",
			Name:                 "Amazing Teacher 1",
			Headline:             "Hi I'm a English Teacher",
			Introduction:         ".........",
			ProfessionalLanguage: []domain.LanguageID{123, 121},
		},
	}

	expectedJSON := types.StringUtil{}.ToRawJSON(`
{
  "data":[
    {
      "id": "1",
      "slug": "at-1",
      "name": "Amazing Teacher 1",
      "headline": "Hi I'm a English Teacher",
      "introduction": ".........",
      "teaching_languages": [123,121],
      "price_info": {
        "trial": 5,
        "normal": 10
      }
    }
  ]
}
`)

	f := new(tutorAndLessonViewFactory)
	actualJSON := f.createArrayData(lessons, tutors)
	ts.Assert().Equal(expectedJSON, string(actualJSON))
}

func (ts *TutorAndLessonViewFactoryTestSuite) Test_createSingleData() {
	lesson := &domain.Lesson{
		LessonID: 2,
		TutorID:  1,
		LessonPrice: domain.LessonPrice{
			Trial:  5.5,
			Normal: 10.5,
		},
	}

	tutor := &domain.Tutor{
		ID:                   1,
		Slug:                 "at-1",
		Name:                 "Amazing Teacher 1",
		Headline:             "Hi I'm a English Teacher",
		Introduction:         ".........",
		ProfessionalLanguage: []domain.LanguageID{123, 121},
	}

	expectedJSON := types.StringUtil{}.ToRawJSON(`
{
  "data": {
    "id": "1",
    "slug": "at-1",
    "name": "Amazing Teacher 1",
    "headline": "Hi I'm a English Teacher",
    "introduction": ".........",
    "teaching_languages": [123,121],
    "price_info": {
      "trial": 5.5,
      "normal": 10.5
    }
  }
}
`)

	f := new(tutorAndLessonViewFactory)
	actualJSON := f.createSingleData(lesson, tutor)
	ts.Assert().Equal(expectedJSON, string(actualJSON))
}
