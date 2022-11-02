package dausach

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	businessError "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	tacgia "daijoubuteam.xyz/se214-library-management/usecase/tac_gia"
	theloai "daijoubuteam.xyz/se214-library-management/usecase/the_loai"
)

type DauSachService struct {
	dauSachRepo    repository.DauSachRepository
	tacGiaUsecase  tacgia.TacGiaUsecase
	theLoaiUsecase theloai.TheLoaiUsecase
}

func (service *DauSachService) GetDanhSachDauSach() ([]*entity.DauSach, error) {
	danhSachDauSach, err := service.dauSachRepo.GetDanhSachDauSach()

	if err != nil {
		return nil, err
	}

	return danhSachDauSach, nil
}

func (service *DauSachService) GetDauSach(maDauSach *entity.ID) (*entity.DauSach, error) {
	dauSach, err := service.dauSachRepo.GetDauSach(maDauSach)
	if err != nil {
		return nil, err
	}

	if dauSach == nil {
		return nil, businessError.NewBusinessError("Dau sach not found")
	}

	return dauSach, nil
}

func (service *DauSachService) CreateDauSach(tenDauSach string, maTheLoai *entity.ID, maTacGia []*entity.ID) (*entity.DauSach, error) {
	tacGia := make([]*entity.TacGia, len(maTacGia))

	for _, mtg := range maTacGia {
		tg, err := service.tacGiaUsecase.GetTacGia(mtg)
		if err != nil {
			return nil, err
		}
		tacGia = append(tacGia, tg)
	}

	theLoai, err := service.theLoaiUsecase.GetTheLoai(maTheLoai)

	if err != nil {
		return nil, err
	}

	dauSach := entity.NewDauSach(theLoai, tenDauSach, tacGia)

	dauSach, err = service.dauSachRepo.CreateDauSach(dauSach)

	if err != nil {
		return nil, err
	}

	return dauSach, nil
}

func (service *DauSachService) UpdateDauSach(maDauSach *entity.ID, tenDauSach string, maTheLoai *entity.ID, maTacGia []*entity.ID) (*entity.DauSach, error) {

	tacGia := make([]*entity.TacGia, len(maTacGia))

	for _, mtg := range maTacGia {
		tg, err := service.tacGiaUsecase.GetTacGia(mtg)
		if err != nil {
			return nil, err
		}
		tacGia = append(tacGia, tg)
	}

	theLoai, err := service.theLoaiUsecase.GetTheLoai(maTheLoai)

	if err != nil {
		return nil, err
	}

	dauSach, err := service.dauSachRepo.GetDauSach(maDauSach)

	if err != nil {
		return nil, err
	}

	dauSach.TenDauSach = tenDauSach
	dauSach.TacGia = tacGia
	dauSach.TheLoai = theLoai

	dauSach, err = service.dauSachRepo.UpdateDauSach(dauSach)

	if err != nil {
		return nil, err
	}

	return dauSach, nil
}

func (service *DauSachService) RemoveDauSach(maDauSach *entity.ID) error {
	err := service.dauSachRepo.RemoveDauSach(maDauSach)
	return err
}
