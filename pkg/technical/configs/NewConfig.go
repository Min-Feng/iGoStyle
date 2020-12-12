package configs

import (
	"os"

	"github.com/rs/zerolog/log"
)

func NewConfig() *Config {
	cfgName := os.Getenv("CFG_NAME")
	center := NewLocalConfigCenter(cfgName)
	cfg := center.GetConfig()
	log.Info().Msg("New Project GetConfig successfully")
	return cfg
}
