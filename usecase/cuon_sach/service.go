package cuonsach

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	nhapsach "daijoubuteam.xyz/se214-library-management/usecase/nhap_sach"
	"daijoubuteam.xyz/se214-library-management/usecase/sach"

	businessError "daijoubuteam.xyz/se214-library-management/core/error"
)

type CuonSachService struct {
	cuonSachRepo    repository.CuonSachRepository
	sachUsecase     sach.SachUsecase
	nhapSachUsecase nhapsach.NhapSachUsecase
}

func (service *CuonSachService) GetDanhSachCuonSach() ([]*entity.CuonSach, error) {
	danhSachCuonSach, err := service.cuonSachRepo.GetDanhSachCuonSach()

	if err != nil {
		return nil, err
	}

	return danhSachCuonSach, nil
}

func (service *CuonSachService) GetCuonSach(maCuonSach *entity.ID) (*entity.CuonSach, error) {
	cuonSach, err := service.cuonSachRepo.GetCuonSach(maCuonSach)

	if err != nil {
		return nil, err
	}

	if cuonSach == nil {
		return nil, businessError.NewBusinessError("cuon sach not found")
	}

	return cuonSach, nil
}

func (service *CuonSachService) CreateCuonSach(maSach *entity.ID, maCtPhieuNhap *entity.ID, tinhTrang bool) (*entity.CuonSach, error) {
	sach, err := service.sachUsecase.GetSach(maSach)
	if err != nil {
		return nil, err
	}

	ctPhieuNhap, err := service.nhapSachUsecase.GetChiTietPhieuNhap(maCtPhieuNhap)

	if err != nil {
		return nil, err
	}

	cuonSach := entity.NewCuonSach(sach, ctPhieuNhap, tinhTrang)

	cuonSach, err = service.cuonSachRepo.CreatCuonSach(cuonSach)

	if err != nil {
		return nil, err
	}

	return cuonSach, nil
}

func (service *CuonSachService) UpdateCuonSach(maCuonSach *entity.ID, maSach *entity.ID, maCtPhieuNhap *entity.ID, tinhTrang bool) (*entity.CuonSach, error) {
	sach, err := service.sachUsecase.GetSach(maSach)
	if err != nil {
		return nil, err
	}

	ctPhieuNhap, err := service.nhapSachUsecase.GetChiTietPhieuNhap(maCtPhieuNhap)

	if err != nil {
		return nil, err
	}

	cuonSach, err := service.GetCuonSach(maCuonSach)

	if err != nil {
		return nil, err
	}

	cuonSach.Sach = sach
	cuonSach.CTPN = ctPhieuNhap
	cuonSach.TinhTrang = tinhTrang

	cuonSach, err = service.cuonSachRepo.UpdateCuonSach(cuonSach)

	if err != nil {
		return nil, err
	}

	return cuonSach, nil
}

func (service *CuonSachService) RemoveCuonSach(maCuonSach *entity.ID) error {
	err := service.cuonSachRepo.RemoveCuonSach(maCuonSach)
	return err
}
