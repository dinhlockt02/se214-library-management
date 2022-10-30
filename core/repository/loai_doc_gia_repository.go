package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type LoaiDocGiaRepository interface {
	CreateLoaiDocGia(loaiDocGia *entity.LoaiDocGia) (*entity.LoaiDocGia, error)
	GetLoaiDocGia(maLoaiDocGia *entity.ID) (*entity.LoaiDocGia, error)
	GetDanhSachLoaiDocGia() ([]*entity.LoaiDocGia, error)
	UpdateLoaiDocGia(loaiDocGia *entity.LoaiDocGia) (*entity.LoaiDocGia, error)
	RemoveLoaiDocGia(maLoaiDocGia *entity.ID) error
}
