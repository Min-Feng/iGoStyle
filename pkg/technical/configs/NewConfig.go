package configs

import (
	"os"

	"github.com/rs/zerolog/log"
)

func NewConfig() *Config {
	cfgName := os.Getenv("CFG_NAME")
	center := NewLocalConfigCenter(cfgName)
	cfg := center.GetConfig()
	log.Info().Msg("GetConfig successfully")
	return cfg
}
