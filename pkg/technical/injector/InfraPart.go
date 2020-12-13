package injector

import (
	"github.com/jmoiron/sqlx"

	"AmazingTalker/pkg/infra/cache"
	"AmazingTalker/pkg/infra/mysql"
	"AmazingTalker/pkg/technical/configs"
)

func InfraPart(cfg *configs.Config) (*sqlx.DB, *cache.Cache) {
	db := mysql.NewMySQL(&cfg.MySQL)
	iCache := cache.NewCache()
	return db, iCache
}
