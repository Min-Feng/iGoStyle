package main

import (
	"github.com/rs/zerolog/log"

	"iGoStyle/pkg/technical/configs"
	"iGoStyle/pkg/technical/injector"
	"iGoStyle/pkg/technical/logger"
)

func init() {
	logger.InfoMode()
}

func main() {
	cfg := configs.NewConfig()
	logger.SetGlobal(cfg.LogLevel, logger.WriterKindHuman)

	router := injector.Server(cfg)

	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Error().Msgf("%v", err)
		if log.Debug().Enabled() {
			log.Error().Msgf("router.Run: ErrorStack=\n%+v", err)
		}
	}
}
