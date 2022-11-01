package theloai

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	businessError "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
)

type TheLoaiService struct {
	theLoaiRepo repository.TheLoaiRepository
}

func (service *TheLoaiService) GetDanhSachTheLoai() ([]*entity.TheLoai, error) {
	danhSachTheLoai, err := service.theLoaiRepo.GetDanhSachTheLoai()

	if err != nil {
		return nil, err
	}

	return danhSachTheLoai, nil
}

func (service *TheLoaiService) GetTheLoai(maTheLoai *entity.ID) (*entity.TheLoai, error) {
	theLoai, err := service.theLoaiRepo.GetTheLoai(maTheLoai)

	if err != nil {
		return nil, err
	}

	if theLoai == nil {
		return nil, businessError.NewBusinessError("The loai not found")
	}

	return theLoai, nil
}

func (service *TheLoaiService) CreateTheLoai(tenTheLoai string) (*entity.TheLoai, error) {
	newTheLoai := entity.NewTheLoai(tenTheLoai)
	if !newTheLoai.IsValid() {
		return nil, businessError.NewBusinessError("Invalid the loai")
	}
	newTheLoai, err := service.theLoaiRepo.CreateTheLoai(newTheLoai)

	if err != nil {
		return nil, err
	}

	return newTheLoai, err
}

func (service *TheLoaiService) UpdateTheLoai(maTheLoai *entity.ID, tenTheLoai string) (*entity.TheLoai, error) {
	updatedTheLoai, err := service.theLoaiRepo.GetTheLoai(maTheLoai)
	if err != nil {
		return nil, err
	}
	if updatedTheLoai == nil {
		return nil, businessError.NewBusinessError("The loai not found")
	}
	updatedTheLoai.TenTheLoai = tenTheLoai
	if !updatedTheLoai.IsValid() {
		return nil, businessError.NewBusinessError("Invalid the loai")
	}
	return updatedTheLoai, nil
}

func (service *TheLoaiService) RemoveTheLoai(maTheLoai *entity.ID) error {
	err := service.theLoaiRepo.RemoveTheLoai(maTheLoai)
	return err
}
