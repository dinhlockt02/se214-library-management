//go:build wireinject
// +build wireinject

package wireimpl

import "github.com/google/wire"

import (
	"daijoubuteam.xyz/se214-library-management/config"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	coreservice "daijoubuteam.xyz/se214-library-management/core/service"
	"daijoubuteam.xyz/se214-library-management/infrastructure/mysql"
	"daijoubuteam.xyz/se214-library-management/infrastructure/service"
	"daijoubuteam.xyz/se214-library-management/usecase/auth"
	thuthu "daijoubuteam.xyz/se214-library-management/usecase/thu_thu"
	"github.com/jmoiron/sqlx"
)

var PasswordHasherSet = wire.NewSet(wire.Bind(new(coreservice.PasswordHasher), new(*service.BcryptPasswordHasher)), service.NewBcryptPasswordHasher)
var JwtTokenServiceSet = wire.NewSet(wire.Bind(new(coreservice.JwtTokenService), new(*service.JwtTokenServiceImpl)), service.NewJwtTokenServiceImpl, config.GetJwtConfig)
var ThuThuRepositorySet = wire.NewSet(wire.Bind(new(repository.ThuThuRepository), new(*mysql.ThuThuRepository)), mysql.NewThuThuRepository)
var ThamSoRepositorySet = wire.NewSet(wire.Bind(new(repository.ThamSoRepository), new(*mysql.ThamSoRepository)), mysql.NewThamSoRepository)
var ThuThuUsecaseSet = wire.NewSet(wire.Bind(new(thuthu.ThuThuUsecase), new(*thuthu.ThuThuService)), thuthu.NewThuThuService, PasswordHasherSet, ThuThuRepositorySet, ThamSoRepositorySet)
var AuthUsecaseSet = wire.NewSet(wire.Bind(new(auth.AuthUsecase), new(*auth.AuthService)), auth.NewAuthService, ThuThuUsecaseSet, JwtTokenServiceSet)

func InitThuThuUsecase(db *sqlx.DB) thuthu.ThuThuUsecase {
	wire.Build(ThuThuUsecaseSet)
	return &thuthu.ThuThuService{}
}

func InitAuthUsecase(db *sqlx.DB) auth.AuthUsecase {
	wire.Build(AuthUsecaseSet)
	return &auth.AuthService{}
}
