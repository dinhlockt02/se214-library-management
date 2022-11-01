package tacgia

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	businessError "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
)

type TacGiaService struct {
	tacGiaRepo repository.TacGiaRepository
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

	if tacGia != nil {
		return nil, businessError.NewBusinessError("Tac gia not found")
	}

	return tacGia, nil
}

func (service *TacGiaService) CreateTacGia(tenTacGia string) (*entity.TacGia, error) {
	newTacGia := entity.NewTacGia(tenTacGia)
	if !newTacGia.IsValid() {
		return nil, businessError.NewBusinessError("Invalid tac gia")
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
		return nil, businessError.NewBusinessError("Tac gia not found")
	}
	updatedTacGia.TenTacGia = tenTacGia
	if !updatedTacGia.IsValid() {
		return nil, businessError.NewBusinessError("Invalid tac gia")
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
