package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"AmazingTalker/pkg/application"
)

func NewTutorAndLessonHandler(uc *application.TutorAndLessonUseCase) *TutorAndLessonHandler {
	return &TutorAndLessonHandler{uc: uc}
}

type TutorAndLessonHandler struct {
	uc *application.TutorAndLessonUseCase
}

func (h TutorAndLessonHandler) GetTutors(ctx *gin.Context) {
	languageSlug := ctx.Param("languageSlug")
	b, err := h.uc.QueryByLanguageSlug(languageSlug)
	if err != nil {
		log.Error().Msgf("%v", err)
		if log.Debug().Enabled() {
			log.Error().Msgf("GetTutors: ErrorStack=\n%+v", err)
		}
		ctx.Data(http.StatusInternalServerError, jsonContextType, nil)
		return
	}
	ctx.Data(http.StatusOK, jsonContextType, b)
}

func (h TutorAndLessonHandler) GetTutor(ctx *gin.Context) {
	tutorSlug := ctx.Param("tutorSlug")
	b, err := h.uc.QueryByTutorSlug(tutorSlug)
	if err != nil {
		log.Error().Msgf("%v", err)
		if log.Debug().Enabled() {
			log.Error().Msgf("GetTutor: ErrorStack=\n%+v", err)
		}
		ctx.Data(http.StatusInternalServerError, jsonContextType, nil)
		return
	}
	ctx.Data(http.StatusOK, jsonContextType, b)
}
