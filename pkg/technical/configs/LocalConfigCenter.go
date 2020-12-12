package configs

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func NewLocalConfigCenter(configFileName string) ConfigCenter {
	if configFileName == "" {
		log.Fatal().Msg("Not found: configFileName is empty")
	}

	workDir, err := os.Getwd()
	workDir = filepath.ToSlash(workDir) // for window os
	if err != nil {
		log.Fatal().Msgf("Not get work directory: %v", err)
	}

	configPath := ""
	for _, dir := range ModuleDirectory {
		if strings.Contains(workDir, dir) {
			path := strings.Split(workDir, dir)
			configPath = path[0] + dir + "/config"
			break
		}
	}

	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName(configFileName)
	vp.AddConfigPath(configPath)
	if err := vp.ReadInConfig(); err != nil {
		log.Fatal().Msgf("Reading config: %v", err)
	}

	log.Info().Msgf("New local config center from %v successfully", vp.ConfigFileUsed())
	return &LocalConfigCenter{vp}
}

type LocalConfigCenter struct {
	*viper.Viper
}

func (center *LocalConfigCenter) GetConfig() *Config {
	cfg := new(Config)

	option := func(c *mapstructure.DecoderConfig) { c.TagName = "configs" }
	if err := center.Viper.Unmarshal(cfg, option); err != nil {
		log.Fatal().Err(err).Msg("Unmarshal Config failed:")
		return nil
	}

	return cfg
}
