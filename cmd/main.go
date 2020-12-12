package main

import (
	"github.com/davecgh/go-spew/spew"

	"AmazingTalker/pkg/technical/configs"
	"AmazingTalker/pkg/technical/logger"
)

func init() {
	logger.InfoMode()
}

func main() {
	cfg := configs.NewConfig()
	spew.Dump(cfg)
	logger.SetGlobal(cfg.LogLevel, logger.WriterKindHuman)
}
