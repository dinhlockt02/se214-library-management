package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type CuonSachRepository interface {
	GetDanhSachCuonSach() ([]*entity.CuonSach, error)
	GetCuonSach(maCuonSach *entity.ID) (*entity.CuonSach, error)
	CreatCuonSach(cuonSach *entity.CuonSach) (*entity.CuonSach, error)
	UpdateCuonSach(*entity.CuonSach) (*entity.CuonSach, error)
	RemoveCuonSach(maCuonSach *entity.ID) error
}
