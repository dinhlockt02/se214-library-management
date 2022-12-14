package theloai

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
)

type TheLoaiService struct {
	theLoaiRepo repository.TheLoaiRepository
}

func NewTheLoaiService(theLoaiRepo repository.TheLoaiRepository) *TheLoaiService {
	return &TheLoaiService{
		theLoaiRepo: theLoaiRepo,
	}
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
		return nil, coreerror.NewNotFoundError("The loai not found", nil)
	}

	return theLoai, nil
}

func (service *TheLoaiService) CreateTheLoai(tenTheLoai string) (*entity.TheLoai, error) {
	newTheLoai := entity.NewTheLoai(tenTheLoai)
	if !newTheLoai.IsValid() {
		return nil, coreerror.NewBadRequestError("Invalid the loai", nil)
	}
	newTheLoai, err := service.theLoaiRepo.CreateTheLoai(newTheLoai)

	if err != nil {
		return nil, err
	}

	return newTheLoai, err
}

func (service *TheLoaiService) UpdateTheLoai(maTheLoai *entity.ID, tenTheLoai string) (*entity.TheLoai, error) {
	updatedTheLoai, err := service.GetTheLoai(maTheLoai)
	if err != nil {
		return nil, err
	}
	if updatedTheLoai == nil {
		return nil, coreerror.NewNotFoundError("The loai not found", nil)
	}
	updatedTheLoai.TenTheLoai = tenTheLoai
	if !updatedTheLoai.IsValid() {
		return nil, coreerror.NewBadRequestError("Invalid the loai", nil)
	}
	updatedTheLoai, err = service.theLoaiRepo.UpdateTheLoai(updatedTheLoai)
	return updatedTheLoai, err
}

func (service *TheLoaiService) RemoveTheLoai(maTheLoai *entity.ID) error {
	err := service.theLoaiRepo.RemoveTheLoai(maTheLoai)
	return err
}
