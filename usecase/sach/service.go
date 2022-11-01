package sach

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	businessError "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	dausach "daijoubuteam.xyz/se214-library-management/usecase/dau_sach"
)

type SachService struct {
	sachRepo       repository.SachRepository
	dauSachUsecase dausach.DauSachUsecase
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
		return nil, businessError.NewBusinessError("Sach not found")
	}

	return sach, nil
}

func (service *SachService) CreateSach(maDauSach *entity.ID, nhaXuatBan string, soLuong uint, triGia uint, namXuatBan uint) (*entity.Sach, error) {
	_, err := service.dauSachUsecase.GetDauSach(maDauSach)
	if err != nil {
		return nil, err
	}

	newSach := entity.NewSach(maDauSach, nhaXuatBan, soLuong, triGia, namXuatBan)

	if !newSach.IsValid() {
		return nil, businessError.NewBusinessError("Invalid sach")
	}

	newSach, err = service.sachRepo.CreateSach(newSach)

	if err != nil {
		return nil, err
	}

	return newSach, nil

}

func (service *SachService) UpdateSach(maSach *entity.ID, maDauSach *entity.ID, nhaXuatBan string, soLuong uint, triGia uint, namXuatBan uint) (*entity.Sach, error) {
	_, err := service.dauSachUsecase.GetDauSach(maDauSach)

	if err != nil {
		return nil, err
	}

	sach, err := service.sachRepo.GetSach(maSach)
	if err != nil {
		return nil, err
	}

	if sach == nil {
		return nil, businessError.NewBusinessError("Sach not found")
	}

	sach.MaDauSach = maDauSach
	sach.NhaXuatBan = nhaXuatBan
	sach.SoLuong = soLuong
	sach.TriGia = triGia
	sach.NamXuatBan = namXuatBan

	if !sach.IsValid() {
		return nil, businessError.NewBusinessError("Invalid sach")
	}

	sach, err = service.sachRepo.UpdateSach(sach)

	if err != nil {
		return nil, err
	}

	return sach, nil
}

func (service *SachService) RemoveSach(maSach *entity.ID) error {
	err := service.sachRepo.RemoveSach(maSach)
	return err
}
