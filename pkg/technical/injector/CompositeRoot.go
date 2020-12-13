package injector

import (
	"github.com/jmoiron/sqlx"

	"AmazingTalker/pkg/application"
	"AmazingTalker/pkg/domain"
	"AmazingTalker/pkg/driving/api"
	"AmazingTalker/pkg/infra/cache"
	"AmazingTalker/pkg/infra/mysql"
	"AmazingTalker/pkg/infra/repository"
	"AmazingTalker/pkg/technical/configs"
)

func Project(cfg *configs.Config) *api.Router {
	db, iCache := InfraPart(cfg)

	tutorRepo, lessonRepo, languageMapTableRepo := Infra(db, iCache)
	tutorAndLessonUseCase := Application(tutorRepo, lessonRepo, languageMapTableRepo)
	router := DrivingAdapter(cfg, tutorAndLessonUseCase)

	return router
}

func Infra(db *sqlx.DB, iCache *cache.Cache) (
	*repository.TutorRepo,
	*repository.LessonRepo,
	*mysql.LanguageMapTableRepo,
) {

	tutorRepo := repository.NewTutorRepo(
		cache.NewTutorRepo(iCache),
		mysql.NewTutorRepo(db),
	)

	lessonRepo := repository.NewLessonRepo(
		cache.NewLessonRepo(iCache),
		mysql.NewLessonRepo(db),
	)

	languageMapTableRepo := mysql.NewLanguageMapTableRepo(db)

	return tutorRepo, lessonRepo, languageMapTableRepo
}

func Application(
	tutorRepo domain.TutorRepo,
	lessonRepo domain.LessonRepo,
	langTableRepo domain.LanguageMapTableRepo,
) *application.TutorAndLessonUseCase {

	tutorAndLessonUC := application.NewTutorAndLessonUseCase(tutorRepo, lessonRepo, langTableRepo)

	return tutorAndLessonUC
}

func DrivingAdapter(cfg *configs.Config, uc *application.TutorAndLessonUseCase) *api.Router {
	router := api.NewRouter(cfg.LogLevel)
	handler := api.NewTutorAndLessonHandler(uc)
	api.Register(router, handler)
	return router
}
