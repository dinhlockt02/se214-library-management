package tra_sach

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"time"
)

type Usecase interface {
	GetPhieuTra() ([]*entity.PhieuTra, error)
	CreatePhieuTra(maSach *entity.ID, ghiChu string, ngayTra *time.Time) (*entity.PhieuTra, error)
}
