package sach

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	dausach "daijoubuteam.xyz/se214-library-management/usecase/dau_sach"
)

type SachService struct {
	sachRepo       repository.SachRepository
	dauSachUsecase dausach.DauSachUsecase
}

func NewSachService(sachRepo repository.SachRepository, dauSachUsecase dausach.DauSachUsecase) *SachService {
	return &SachService{
		sachRepo:       sachRepo,
		dauSachUsecase: dauSachUsecase,
	}
}

func (service *SachService) GetDanhSachSach() ([]*entity.Sach, error) {
	danhSachSach, err := service.sachRepo.GetDanhSachSach()
	if err != nil {
		return nil, err
	}

	return danhSachSach, nil
}

func (service *SachService) GetSach(maSach *entity.ID) (*entity.Sach, error) {
	sach, err := service.sachRepo.GetSach(maSach)

	if err != nil {
		return nil, err
	}

	if sach == nil {
		return nil, coreerror.NewNotFoundError("Sach not found", nil)
	}

	return sach, nil
}

func (service *SachService) UpdateSach(maSach *entity.ID, nhaXuatBan string, triGia uint, namXuatBan uint, tinhTrang bool, ghiChu string) (*entity.Sach, error) {

	sach, err := service.sachRepo.GetSach(maSach)
	if err != nil {
		return nil, err
	}

	if sach == nil {
		return nil, coreerror.NewNotFoundError("Sach not found", nil)
	}

	sach.NhaXuatBan = nhaXuatBan
	sach.TriGia = triGia
	sach.NamXuatBan = namXuatBan
	sach.GhiChu = ghiChu
	sach.TinhTrang = tinhTrang

	if !sach.IsValid() {
		return nil, coreerror.NewBadRequestError("Invalid sach", nil)
	}

	sach, err = service.sachRepo.UpdateSach(sach)

	if err != nil {
		return nil, err
	}

	return sach, nil
}
