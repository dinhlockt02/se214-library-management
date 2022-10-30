package docgia

import (
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
)

type DocGiaUsecase interface {
	getDanhSachDocGia() ([]*entity.DocGia, error)
	getDocGia(maDocGia *entity.ID) (*entity.DocGia, error)
	createDocGia(hoTen string, loaiDocGia entity.ID, ngaySinh time.Time, diaChi string, email string, ngayLapThe time.Time) (*entity.DocGia, error)
	updateDocGia(maDocGia entity.ID, hoTen string, loaiDocGia entity.ID, ngaySinh time.Time, diaChi string, email string, ngayLapThe time.Time) (*entity.DocGia, error)
	removeDocGia(maDocGia *entity.ID) (*entity.DocGia, error)
}
