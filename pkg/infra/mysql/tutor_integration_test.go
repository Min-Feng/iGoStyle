// +build integration

package mysql_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"

	"AmazingTalker/pkg/infra/mysql"
	"AmazingTalker/pkg/technical/configs"
	"AmazingTalker/pkg/technical/injector"
	"AmazingTalker/pkg/technical/logger"
)

func TestTutorRepo_QueryAllByLanguageID(t *testing.T) {
	logger.DebugMode()
	cfg := configs.NewConfig()
	db, _ := injector.InfraPart(cfg)
	tutorRepo := mysql.NewTutorRepo(db)

	tutors, err := tutorRepo.QueryAllByLanguageID(123)
	assert.NoError(t, err)
	log.Debug().Msgf("\n%v", spew.Sdump(tutors))
}
