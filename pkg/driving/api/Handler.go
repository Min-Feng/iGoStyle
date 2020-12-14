package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"iGoStyle/pkg/application"
)

func NewHandler(uc *application.TutorAndLessonUseCase) *Handler {
	return &Handler{uc: uc}
}

type Handler struct {
	uc *application.TutorAndLessonUseCase
}

func (h Handler) GetTutors(ctx *gin.Context) {
	languageSlug := ctx.Param("languageSlug")
	response, err := h.uc.QueryByLanguageSlug(languageSlug)
	if err != nil {
		log.Error().Msgf("%v", err)
		if log.Debug().Enabled() {
			log.Error().Msgf("GetTutors: ErrorStack=\n%+v", err)
		}
		ctx.Data(http.StatusInternalServerError, jsonContextType, nil)
		return
	}
	ctx.Data(http.StatusOK, jsonContextType, response)
}

func (h Handler) GetTutor(ctx *gin.Context) {
	tutorSlug := ctx.Param("tutorSlug")
	response, err := h.uc.QueryByTutorSlug(tutorSlug)
	if err != nil {
		log.Error().Msgf("%v", err)
		if log.Debug().Enabled() {
			log.Error().Msgf("GetTutor: ErrorStack=\n%+v", err)
		}
		ctx.Data(http.StatusInternalServerError, jsonContextType, nil)
		return
	}
	ctx.Data(http.StatusOK, jsonContextType, response)
}
