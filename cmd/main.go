package main

import (
	"github.com/rs/zerolog/log"

	"AmazingTalker/pkg/technical/configs"
	"AmazingTalker/pkg/technical/injector"
	"AmazingTalker/pkg/technical/logger"
)

func init() {
	logger.InfoMode()
}

func main() {
	cfg := configs.NewConfig()
	logger.SetGlobal(cfg.LogLevel, logger.WriterKindHuman)
	// spew.Dump(cfg)

	router := injector.Project(cfg)

	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Error().Msgf("%v", err)
		if log.Debug().Enabled() {
			log.Error().Msgf("router.Run: ErrorStack=\n%+v", err)
		}
	}
}
