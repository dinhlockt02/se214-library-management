package config

import (
	"time"
)

const (
	Release string = "RELEASE"
	Debug          = "DEBUG"
)

type Config struct {
	Mode           string
	DatabaseConfig DatabaseConfig
	JwtConfig      JwtConfig
}

type JwtConfig struct {
	Audience    string
	Secret      []byte
	Issuer      string
	ExpDuration time.Duration
}

type DatabaseConfig struct {
	Driver   string
	User     string
	Password string
	Host     string
	Name     string
	Port     string
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
