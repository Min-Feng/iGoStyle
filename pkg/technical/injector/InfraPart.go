package injector

import (
	"github.com/jmoiron/sqlx"

	"iGoStyle/pkg/infra/cache"
	"iGoStyle/pkg/infra/mysql"
	"iGoStyle/pkg/technical/configs"
)

func InfraPart(cfg *configs.Config) (*sqlx.DB, *cache.Cache) {
	db := mysql.NewMySQL(&cfg.MySQL)
	iCache := cache.NewCache()
	return db, iCache
}
