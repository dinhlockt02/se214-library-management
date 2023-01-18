package repository

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
)

type PhieuTraRepository interface {
	GetDanhSachPhieuTra() ([]*entity.PhieuTra, error)
	GetPhieuTraByDocGia(maDocGia *entity.ID) ([]*entity.PhieuTra, error)
	CreatePhieuTra(phieuTra *entity.PhieuTra) (*entity.PhieuTra, error)
}
