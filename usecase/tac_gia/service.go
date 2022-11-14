package tacgia

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
)

type TacGiaService struct {
	tacGiaRepo repository.TacGiaRepository
}

func NewTacGiaService(tacGiaRepo repository.TacGiaRepository) *TacGiaService {
	return &TacGiaService{
		tacGiaRepo: tacGiaRepo,
	}
}

func (service *TacGiaService) GetDanhSachTacGia() ([]*entity.TacGia, error) {
	danhSachTacGia, err := service.tacGiaRepo.GetDanhSachTacGia()
	if err != nil {
		return nil, err
	}
	return danhSachTacGia, nil
}

func (service *TacGiaService) GetTacGia(maTacGia *entity.ID) (*entity.TacGia, error) {
	tacGia, err := service.tacGiaRepo.GetTacGia(maTacGia)
	if err != nil {
		return nil, err
	}

	if tacGia == nil {
		return nil, coreerror.NewNotFoundError("Tac gia not found", nil)
	}

	return tacGia, nil
}

func (service *TacGiaService) CreateTacGia(tenTacGia string) (*entity.TacGia, error) {
	newTacGia := entity.NewTacGia(tenTacGia)
	if !newTacGia.IsValid() {
		return nil, coreerror.NewBadRequestError("Invalid tac gia", nil)
	}
	newTacGia, err := service.tacGiaRepo.CreateTacGia(newTacGia)
	if err != nil {
		return nil, err
	}
	return newTacGia, nil
}

func (service *TacGiaService) UpdateTacGia(maTacGia *entity.ID, tenTacGia string) (*entity.TacGia, error) {
	updatedTacGia, err := service.tacGiaRepo.GetTacGia(maTacGia)
	if err != nil {
		return nil, err
	}
	if updatedTacGia == nil {
		return nil, coreerror.NewNotFoundError("Tac gia not found", nil)
	}
	updatedTacGia.TenTacGia = tenTacGia
	if !updatedTacGia.IsValid() {
		return nil, coreerror.NewBadRequestError("Invalid tac gia", nil)
	}
	_, err = service.tacGiaRepo.UpdateTacGia(updatedTacGia)
	if err != nil {
		return nil, err
	}
	return updatedTacGia, nil
}

func (service *TacGiaService) DeleteTacGia(maTacGia *entity.ID) error {
	err := service.tacGiaRepo.RemoveTacGia(maTacGia)
	return err
}
