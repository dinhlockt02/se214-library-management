package repository

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
)

type PhieuThuTienRepository interface {
	GetPhieuThuTienByMaDocGia(maDocGia string) ([]*entity.PhieuThuTien, error)
	CreatePhieuThuTien(phieuThuTien *entity.PhieuThuTien) (*entity.PhieuThuTien, error)
}
