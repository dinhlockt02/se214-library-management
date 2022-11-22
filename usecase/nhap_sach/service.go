package nhapsach

import (
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	dausach "daijoubuteam.xyz/se214-library-management/usecase/dau_sach"
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
	"daijoubuteam.xyz/se214-library-management/core/repository"
)

type NhapSachService struct {
	phieuNhapRepo  repository.PhieuNhapRepository
	dauSachUsecase dausach.DauSachUsecase
}

func NewNhapSachService(phieuNhapRepository repository.PhieuNhapRepository, dauSachUsecase dausach.DauSachUsecase) *NhapSachService {
	return &NhapSachService{
		phieuNhapRepo:  phieuNhapRepository,
		dauSachUsecase: dauSachUsecase,
	}
}

func (service *NhapSachService) GetDanhSachPhieuNhapSach() ([]*entity.PhieuNhap, error) {
	danhSachPhieuNhap, err := service.phieuNhapRepo.GetDanhSachPhieuNhap()

	if err != nil {
		return nil, err
	}

	return danhSachPhieuNhap, nil
}

func (service *NhapSachService) GetPhieuNhapSach(maPhieuNhap *entity.ID) (*entity.PhieuNhap, error) {
	phieuNhap, err := service.phieuNhapRepo.GetPhieuNhap(maPhieuNhap)

	if err != nil {
		return nil, err
	}

	return phieuNhap, nil
}

func (service *NhapSachService) CreatePhieuNhapSach(ngayLap *time.Time) (*entity.PhieuNhap, error) {

	phieuNhap := entity.NewPhieuNhap(ngayLap)

	if !phieuNhap.IsValid() {
		return nil, coreerror.NewBadRequestError("Invalid phieu nhap", nil)
	}

	phieuNhap, err := service.phieuNhapRepo.CreatePhieuNhap(phieuNhap)

	if err != nil {
		return nil, err
	}

	return phieuNhap, nil
}

func (service *NhapSachService) UpdatePhieuNhapSach(maPhieuNhap *entity.ID, ngayLap *time.Time) (*entity.PhieuNhap, error) {
	phieuNhap, err := service.GetPhieuNhapSach(maPhieuNhap)
	if err != nil {
		return nil, err
	}

	phieuNhap.NgayLap = ngayLap

	if !phieuNhap.IsValid() {
		return nil, coreerror.NewBadRequestError("Invalid phieu nhap", nil)

	}
	phieuNhap, err = service.phieuNhapRepo.UpdatePhieuNhap(phieuNhap)

	if err != nil {
		return nil, err
	}

	return phieuNhap, nil
}

func (service *NhapSachService) RemovePhieuNhapSach(maPhieuNhap *entity.ID) error {
	phieuNhap, err := service.GetPhieuNhapSach(maPhieuNhap)
	if err != nil {
		return err
	}

	err = service.phieuNhapRepo.RemovePhieuNhap(phieuNhap)
	return err
}

func (service *NhapSachService) AddChiTietPhieuNhapSach(
	maPhieuNhap *entity.ID,
	maDauSach *entity.ID,
	nhaXuatBan string, triGia uint,
	namXuatBan uint,
	tinhTrang bool,
	donGia uint) (*entity.CtPhieuNhap, error) {
	dauSach, err := service.dauSachUsecase.GetDauSach(maDauSach)
	if err != nil {
		return nil, err
	}
	sach := entity.NewSach(dauSach, nhaXuatBan, triGia, namXuatBan, tinhTrang)
	ctNhapSach := entity.NewCtPhieuNhap(sach, donGia)
	chiTietPhieuNhap, err := service.phieuNhapRepo.AddChiTietPhieuNhap(maPhieuNhap, ctNhapSach)
	if err != nil {
		return nil, err
	}
	return chiTietPhieuNhap, nil
}

func (service *NhapSachService) RemoveChiTietPhieuNhapSach(maChiTietPhieuNhap *entity.ID) error {
	return service.phieuNhapRepo.RemoveChiTietPhieuNhap(maChiTietPhieuNhap)
}
