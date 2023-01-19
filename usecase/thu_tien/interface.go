package thu_tien

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"time"
)

type Usecase interface {
	GetPhieuThuTienByMaDocGia(maDocGia string) ([]*entity.PhieuThuTien, error)
	CreatePhieuThuTien(maDocGia string, soTienThu int, ngayThu *time.Time) (*entity.PhieuThuTien, error)
}
