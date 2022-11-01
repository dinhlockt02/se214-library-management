package loaidocgia

import "daijoubuteam.xyz/se214-library-management/core/entity"

type LoaiDocGiaUsecase interface {
	GetDanhSachLoaiDocGia() ([]*entity.LoaiDocGia, error)
	GetLoaiDocGia(maLoaiDocGia *entity.ID) (*entity.LoaiDocGia, error)
	CreateLoaiDocGia(tenLoaiDocGia string) (*entity.LoaiDocGia, error)
	UpdateLoaiDocGia(maLoaiDocGia *entity.ID, tenLoaiDocGia string) (*entity.LoaiDocGia, error)
	DeleteLoaiDocGia(maLoaiDocGia *entity.ID) error
}
