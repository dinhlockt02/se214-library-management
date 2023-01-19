package docgia

import (
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
)

type DocGiaUsecase interface {
	GetDanhSachDocGia() ([]*entity.DocGia, error)
	GetDocGia(maDocGia string) (*entity.DocGia, error)
	CreateDocGia(maDocGia string, hoTen string, maLoaiDocGia *entity.ID, ngaySinh *time.Time, diaChi string, email string, ngayLapThe *time.Time) (*entity.DocGia, error)
	UpdateDocGia(maDocGia string, hoTen string, maLoaiDocGia *entity.ID, ngaySinh *time.Time, diaChi string, email string, ngayLapThe *time.Time) (*entity.DocGia, error)
	RemoveDocGia(maDocGia string) error
}
