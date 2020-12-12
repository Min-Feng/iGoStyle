package configs

import (
	"fmt"

	"AmazingTalker/pkg/technical/logger"
)

type ConfigCenter interface {
	GetConfig() *Config
}

type Config struct {
	Port     string       `configs:"port"`
	LogLevel logger.Level `configs:"log_level"`
	MySQL    MySQL        `configs:"mysql"`
}

type MySQL struct {
	User        string `configs:"user"`
	Password    string `configs:"password"`
	Host        string `configs:"host"`
	Port        string `configs:"port"`
	Database    string `configs:"database"`
	MaxConn     int    `configs:"max_conn"`
	MaxIdleConn int    `configs:"max_idle_conn"`
}

func (c *MySQL) DSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
}
