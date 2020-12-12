package injector

import (
	"github.com/jmoiron/sqlx"

	"AmazingTalker/pkg/infra/mysql"
	"AmazingTalker/pkg/technical/configs"
)

func InfraPart() *sqlx.DB {
	cfg := configs.NewConfig()
	db := mysql.NewMySQL(&cfg.MySQL)
	return db
}
