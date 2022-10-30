package loaidocgia

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	businessError "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
)

type LoaiDocGiaService struct {
	loaiDocGiaRepo repository.LoaiDocGiaRepository
}

func (service *LoaiDocGiaService) getDanhSachLoaiDocGia() ([]*entity.LoaiDocGia, error) {
	danhSachDocGia, err := service.loaiDocGiaRepo.GetDanhSachLoaiDocGia()

	if err != nil {
		return nil, err
	}

	return danhSachDocGia, nil
}

func (service *LoaiDocGiaService) getLoaiDocGia(maLoaiDocGia *entity.ID) (*entity.LoaiDocGia, error) {
	docGia, err := service.loaiDocGiaRepo.GetLoaiDocGia(maLoaiDocGia)

	if err != nil {
		return nil, err
	}

	if docGia == nil {
		return nil, businessError.NewBusinessError("doc gia not found")
	}

	return docGia, nil
}

func (service *LoaiDocGiaService) createLoaiDocGia(tenLoaiDocGia string) (*entity.LoaiDocGia, error) {
	loaiDocGia := entity.NewLoaiDocGia(tenLoaiDocGia)
	_, err := service.loaiDocGiaRepo.CreateLoaiDocGia(loaiDocGia)
	if err != nil {
		return nil, err
	}

	return loaiDocGia, err
}

func (service *LoaiDocGiaService) updateLoaiDocGia(maLoaiDocGia *entity.ID, tenLoaiDocGia string) (*entity.LoaiDocGia, error) {
	loaiDocGia, err := service.loaiDocGiaRepo.GetLoaiDocGia(maLoaiDocGia)

	if err != nil {
		return nil, err
	}

	if loaiDocGia == nil {
		return nil, businessError.NewBusinessError("loai doc gia not found")
	}

	loaiDocGia.TenLoaiDocGia = tenLoaiDocGia

	_, err = service.loaiDocGiaRepo.UpdateLoaiDocGia(loaiDocGia)

	if err != nil {
		return nil, err
	}

	return loaiDocGia, nil
}

func (service *LoaiDocGiaService) deleteLoaiDocGia(maLoaiDocGia *entity.ID) error {
	err := service.loaiDocGiaRepo.RemoveLoaiDocGia(maLoaiDocGia)
	return err
}
