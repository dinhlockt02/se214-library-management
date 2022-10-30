package loaidocgia

import "daijoubuteam.xyz/se214-library-management/core/entity"

type LoaiDocGiaUsecase interface {
	getDanhSachLoaiDocGia() ([]*entity.LoaiDocGia, error)
	getLoaiDocGia(maLoaiDocGia *entity.ID) (*entity.LoaiDocGia, error)
	createLoaiDocGia(tenLoaiDocGia string) (*entity.LoaiDocGia, error)
	updateLoaiDocGia(maLoaiDocGia *entity.ID, tenLoaiDocGia string) (*entity.LoaiDocGia, error)
	deleteLoaiDocGia(maLoaiDocGia *entity.ID) error
}
