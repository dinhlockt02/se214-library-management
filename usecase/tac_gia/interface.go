package tacgia

import "daijoubuteam.xyz/se214-library-management/core/entity"

type TacGiaUsecase interface {
	GetDanhSachTacGia() ([]*entity.TacGia, error)
	GetTacGia(maTacGia *entity.ID) (*entity.TacGia, error)
	CreateTacGia(tenTacGia string) (*entity.TacGia, error)
	UpdateTacGia(maTacGia *entity.ID, tenTacGia string) (*entity.TacGia, error)
	DeleteTacGia(maTacGia *entity.ID) error
}
