package injector

import (
	"github.com/jmoiron/sqlx"

	"AmazingTalker/pkg/infra/mysql"
)

func Infra(db *sqlx.DB) (*mysql.TutorRepo, *mysql.LessonRepo, *mysql.LanguageMapTableRepo) {
	lessonRepo := mysql.NewLessonRepo(db)
	tutorRepo := mysql.NewTutorRepo(db)
	languageMapTableRepo := mysql.NewLanguageMapTableRepo(db)
	return tutorRepo, lessonRepo, languageMapTableRepo
}
