package thuthucommandlogin

import (
	"daijoubuteam.xyz/se214-library-management/config"
	"daijoubuteam.xyz/se214-library-management/infrastructure/mysql"
	"daijoubuteam.xyz/se214-library-management/infrastructure/service"
	"daijoubuteam.xyz/se214-library-management/usecase/auth"
	thuthu "daijoubuteam.xyz/se214-library-management/usecase/thu_thu"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Login(db *sqlx.DB, email string, password string) {

	thuThuRepo := mysql.NewThuThuRepository(db)
	thamSoRepo := mysql.NewThamSoRepository(db)
	passwordHasher := service.NewBcryptPasswordHasher()
	jwtTokenService := service.NewJwtTokenServiceImpl(config.DevConfig.JwtConfig)
	thuThuService := thuthu.NewThuThuService(passwordHasher, thuThuRepo, thamSoRepo)
	authSerivce := auth.NewAuthService(thuThuService, passwordHasher, jwtTokenService)

	token, err := authSerivce.Login(email, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("User token is: %v\n", *token)
}
