package config

import "time"

var TestConfig = &Config{
	DbName:         "se214",
	User:           "root",
	Password:       "example",
	Host:           "db",
	DatabaseDriver: "mysql",
	Port:           "3306",
	JwtConfig: JwtConfig{
		Audience:    "localhost",
		Secret:      []byte("localhost"),
		Issuer:      "localhost",
		ExpDuration: time.Duration(24) * time.Hour,
	},
}
