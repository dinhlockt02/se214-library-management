package loaidocgia

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
)

type LoaiDocGiaService struct {
	loaiDocGiaRepo repository.LoaiDocGiaRepository
}

func NewLoaiDocGiaService(loaiDocGiaRepo repository.LoaiDocGiaRepository) *LoaiDocGiaService {
	return &LoaiDocGiaService{
		loaiDocGiaRepo: loaiDocGiaRepo,
	}
}

func (service *LoaiDocGiaService) GetDanhSachLoaiDocGia() ([]*entity.LoaiDocGia, error) {
	danhSachDocGia, err := service.loaiDocGiaRepo.GetDanhSachLoaiDocGia()

	if err != nil {
		return nil, err
	}

	return danhSachDocGia, nil
}

func (service *LoaiDocGiaService) GetLoaiDocGia(maLoaiDocGia *entity.ID) (*entity.LoaiDocGia, error) {
	loaiDocGia, err := service.loaiDocGiaRepo.GetLoaiDocGia(maLoaiDocGia)

	if err != nil {
		return nil, err
	}

	if loaiDocGia == nil {
		return nil, coreerror.NewNotFoundError("loai doc gia not found", nil)
	}

	return loaiDocGia, nil
}

func (service *LoaiDocGiaService) CreateLoaiDocGia(tenLoaiDocGia string) (*entity.LoaiDocGia, error) {
	loaiDocGia := entity.NewLoaiDocGia(tenLoaiDocGia)
	_, err := service.loaiDocGiaRepo.CreateLoaiDocGia(loaiDocGia)
	if err != nil {
		return nil, err
	}

	return loaiDocGia, err
}

func (service *LoaiDocGiaService) UpdateLoaiDocGia(maLoaiDocGia *entity.ID, tenLoaiDocGia string) (*entity.LoaiDocGia, error) {
	loaiDocGia, err := service.loaiDocGiaRepo.GetLoaiDocGia(maLoaiDocGia)

	if err != nil {
		return nil, err
	}

	if loaiDocGia == nil {
		return nil, coreerror.NewNotFoundError("loai doc gia not found", nil)
	}

	loaiDocGia.TenLoaiDocGia = tenLoaiDocGia

	_, err = service.loaiDocGiaRepo.UpdateLoaiDocGia(loaiDocGia)

	if err != nil {
		return nil, err
	}

	return loaiDocGia, nil
}

func (service *LoaiDocGiaService) DeleteLoaiDocGia(maLoaiDocGia *entity.ID) error {
	err := service.loaiDocGiaRepo.RemoveLoaiDocGia(maLoaiDocGia)
	return err
}
