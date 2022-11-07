package config

import "time"

type Config struct {
	DatabaseDriver string
	User           string
	Password       string
	Host           string
	DbName         string
	Port           string
	JwtConfig      JwtConfig
}

type JwtConfig struct {
	Audience    string
	Secret      []byte
	Issuer      string
	ExpDuration time.Duration
}

func GetConfig() *Config {
	return DevConfig
}

func GetJwtConfig() JwtConfig {
	return GetConfig().JwtConfig
}
