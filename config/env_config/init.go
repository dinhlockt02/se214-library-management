package env_config

import (
	"daijoubuteam.xyz/se214-library-management/config"
	"os"
	"strconv"
	"time"
)

func init() {
	DatabaseDriver, exists := os.LookupEnv("DATABASE_DRIVER")
	DatabaseUser, exists := os.LookupEnv("DATABASE_USER")
	DatabasePassword, exists := os.LookupEnv("DATABASE_PASSWORD")
	DatabaseHost, exists := os.LookupEnv("DATABASE_HOST")
	DatabaseName, exists := os.LookupEnv("DATABASE_NAME")
	DatabasePort, exists := os.LookupEnv("DATABASE_PORT")

	JwtAudience, exists := os.LookupEnv("JWT_AUDIENCE")
	JwtSecret, exists := os.LookupEnv("JWT_SECRET")
	JwtIssuer, exists := os.LookupEnv("JWT_ISSUER")
	JWTExpDurationStr, exists := os.LookupEnv("JWT_EXP_DURATION")

	if !exists {
		panic("environment variable missing")
	}

	JWTExpDuration, err := strconv.Atoi(JWTExpDurationStr)

	if err != nil {
		panic(err)
	}

	config.SetConfig(config.Config{
		DatabaseDriver:   DatabaseDriver,
		DatabaseUser:     DatabaseUser,
		DatabasePassword: DatabasePassword,
		DatabaseHost:     DatabaseHost,
		DatabasePort:     DatabasePort,
		DatabaseName:     DatabaseName,
		JwtConfig: config.JwtConfig{
			Audience:    JwtAudience,
			Secret:      []byte(JwtSecret),
			Issuer:      JwtIssuer,
			ExpDuration: time.Duration(JWTExpDuration) * time.Hour,
		},
	})
}
