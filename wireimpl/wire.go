//go:build wireinject
// +build wireinject

package wireimpl

import (
	dausach "daijoubuteam.xyz/se214-library-management/usecase/dau_sach"
	docgia "daijoubuteam.xyz/se214-library-management/usecase/doc_gia"
	loaidocgia "daijoubuteam.xyz/se214-library-management/usecase/loai_doc_gia"
	nhapsach "daijoubuteam.xyz/se214-library-management/usecase/nhap_sach"
	tacgia "daijoubuteam.xyz/se214-library-management/usecase/tac_gia"
	theloai "daijoubuteam.xyz/se214-library-management/usecase/the_loai"
	"github.com/google/wire"
)

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
var LoaiDocGiaRepositorySet = wire.NewSet(wire.Bind(new(repository.LoaiDocGiaRepository), new(*mysql.LoaiDocGiaRepository)), mysql.NewLoaiDocGiaRepository)
var DocGiaRepositorySet = wire.NewSet(wire.Bind(new(repository.DocGiaRepository), new(*mysql.DocGiaRepository)), mysql.NewDocGiaRepository)
var TheLoaiRepositorySet = wire.NewSet(wire.Bind(new(repository.TheLoaiRepository), new(*mysql.TheLoaiRepository)), mysql.NewTheLoaiRepository)
var TacGiaRepositorySet = wire.NewSet(wire.Bind(new(repository.TacGiaRepository), new(*mysql.TacGiaRepository)), mysql.NewTacGiaRepository)
var DauSachRepositorySet = wire.NewSet(wire.Bind(new(repository.DauSachRepository), new(*mysql.DauSachRepository)), mysql.NewDauSachRepository)
var PhieuNhapRepositorySet = wire.NewSet(wire.Bind(new(repository.PhieuNhapRepository), new(*mysql.PhieuNhapRepository)), mysql.NewPhieuNhapRepository)

var ThuThuUsecaseSet = wire.NewSet(wire.Bind(new(thuthu.ThuThuUsecase), new(*thuthu.ThuThuService)), thuthu.NewThuThuService, PasswordHasherSet, ThuThuRepositorySet, ThamSoRepositorySet)
var AuthUsecaseSet = wire.NewSet(wire.Bind(new(auth.AuthUsecase), new(*auth.AuthService)), auth.NewAuthService, ThuThuUsecaseSet, JwtTokenServiceSet)
var LoaiDocGiaUsecaseSet = wire.NewSet(wire.Bind(new(loaidocgia.LoaiDocGiaUsecase), new(*loaidocgia.LoaiDocGiaService)), loaidocgia.NewLoaiDocGiaService, LoaiDocGiaRepositorySet)
var DocGiaUsecaseSet = wire.NewSet(wire.Bind(new(docgia.DocGiaUsecase), new(*docgia.DocGiaService)), docgia.NewDocGiaService, DocGiaRepositorySet, LoaiDocGiaUsecaseSet, ThamSoRepositorySet)
var TheLoaiUsecaseSet = wire.NewSet(wire.Bind(new(theloai.TheLoaiUsecase), new(*theloai.TheLoaiService)), theloai.NewTheLoaiService, TheLoaiRepositorySet)
var TacGiaUsecaseSet = wire.NewSet(wire.Bind(new(tacgia.TacGiaUsecase), new(*tacgia.TacGiaService)), tacgia.NewTacGiaService, TacGiaRepositorySet)
var DauSachUsecaseSet = wire.NewSet(wire.Bind(new(dausach.DauSachUsecase), new(*dausach.DauSachService)), dausach.NewDauSachService, TacGiaUsecaseSet, TheLoaiUsecaseSet, DauSachRepositorySet)
var NhapSachUsecaseSet = wire.NewSet(wire.Bind(new(nhapsach.NhapSachUsecase), new(*nhapsach.NhapSachService)), nhapsach.NewNhapSachService, DauSachUsecaseSet, PhieuNhapRepositorySet)

func InitThuThuUsecase(db *sqlx.DB) thuthu.ThuThuUsecase {
	wire.Build(ThuThuUsecaseSet)
	return &thuthu.ThuThuService{}
}

func InitAuthUsecase(db *sqlx.DB) auth.AuthUsecase {
	wire.Build(AuthUsecaseSet)
	return &auth.AuthService{}
}

func InitLoaiDocGiaUsecase(db *sqlx.DB) loaidocgia.LoaiDocGiaUsecase {
	wire.Build(LoaiDocGiaUsecaseSet)
	return &loaidocgia.LoaiDocGiaService{}
}

func InitDocGiaUsecase(db *sqlx.DB) docgia.DocGiaUsecase {
	wire.Build(DocGiaUsecaseSet)
	return &docgia.DocGiaService{}
}

func InitTheLoaiUsecase(db *sqlx.DB) theloai.TheLoaiUsecase {
	wire.Build(TheLoaiUsecaseSet)
	return &theloai.TheLoaiService{}
}

func InitTacGiaUsecase(db *sqlx.DB) tacgia.TacGiaUsecase {
	wire.Build(TacGiaUsecaseSet)
	return &tacgia.TacGiaService{}
}

func InitDauSachUsecase(db *sqlx.DB) dausach.DauSachUsecase {
	wire.Build(DauSachUsecaseSet)
	return &dausach.DauSachService{}
}

func InitNhapSachUsecase(db *sqlx.DB) nhapsach.NhapSachUsecase {
	wire.Build(NhapSachUsecaseSet)
	return &nhapsach.NhapSachService{}
}
