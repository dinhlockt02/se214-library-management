package cuonsach

import "daijoubuteam.xyz/se214-library-management/core/entity"

type CuonSachUsecase interface {
	GetDanhSachCuonSach() ([]*entity.CuonSach, error)
	GetCuonSach(maCuonSach *entity.ID) (*entity.CuonSach, error)
	CreateCuonSach(maSach *entity.ID, maCtPhieuNhap *entity.ID, tinhTrang bool) (*entity.CuonSach, error)
	UpdateCuonSach(maCuonSach *entity.ID, maSach *entity.ID, maCtPhieuNhap *entity.ID, tinhTrang bool) (*entity.CuonSach, error)
	RemoveCuonSach(maCuonSach *entity.ID) error
}
