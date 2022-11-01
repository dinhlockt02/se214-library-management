package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type TacGiaRepository interface {
	CreateTacGia(tacGia *entity.TacGia) (*entity.TacGia, error)
	GetTacGia(maTacGia *entity.ID) (*entity.TacGia, error)
	GetDanhSachTacGia() ([]*entity.TacGia, error)
	UpdateTacGia(tacGia *entity.TacGia) (*entity.TacGia, error)
	RemoveTacGia(maTacGia *entity.ID) error
}
