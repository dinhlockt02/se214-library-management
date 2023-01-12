package config

import (
	"time"
)

type Config struct {
	DatabaseDriver   string
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabaseName     string
	DatabasePort     string
	JwtConfig        JwtConfig
}

type JwtConfig struct {
	Audience    string
	Secret      []byte
	Issuer      string
	ExpDuration time.Duration
}

var config Config

func GetConfig() Config {
	return config
}

func GetJwtConfig() JwtConfig {
	return GetConfig().JwtConfig
}

func SetConfig(cf Config) {
	config = cf
}
