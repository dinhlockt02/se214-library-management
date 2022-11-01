package theloai

import "daijoubuteam.xyz/se214-library-management/core/entity"

type TheLoaiUsecase interface {
	GetDanhSachTheLoai() ([]*entity.TheLoai, error)
	GetTheLoai(maTheLoai *entity.ID) (*entity.TheLoai, error)
	CreateTheLoai(tenTheLoai string) (*entity.TheLoai, error)
	UpdateTheLoai(maTheLoai *entity.ID, tenTheLoai string) (*entity.TheLoai, error)
	RemoveTheLoai(maTheLoai *entity.ID) error
}
