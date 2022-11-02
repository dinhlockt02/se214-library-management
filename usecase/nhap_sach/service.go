package nhapsach

import (
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	"daijoubuteam.xyz/se214-library-management/usecase/sach"

	businessError "daijoubuteam.xyz/se214-library-management/core/error"
)

type NhapSachService struct {
	phieuNhapRepo repository.PhieuNhapRepository
	sachUsecase   sach.SachUsecase
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
		return nil, businessError.NewBusinessError("Invalid phieu nhap")
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
		return nil, businessError.NewBusinessError("Invalid phieu nhap")

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

func (service *NhapSachService) AddChiTietPhieuNhapSach(maPhieuNhap *entity.ID, maSach *entity.ID, soLuong uint, donGia uint) (*entity.PhieuNhap, error) {

	sach, err := service.sachUsecase.GetSach(maSach)

	if err != nil {
		return nil, err
	}

	phieuNhap, err := service.GetPhieuNhapSach(maPhieuNhap)

	if err != nil {
		return nil, err
	}

	ctPhieuNhap := entity.NewCtPhieuNhap(phieuNhap, sach, soLuong, donGia)

	if !ctPhieuNhap.IsValid() {
		return nil, businessError.NewBusinessError("Invalid chi tiet phieu nhap")
	}

	phieuNhap, err = service.phieuNhapRepo.AddChiTietPhieuNhap(ctPhieuNhap)

	if err != nil {
		return nil, err
	}

	return phieuNhap, err
}

func (service *NhapSachService) RemoveChiTietPhieuNhapSach(maChiTietPhieuNhap *entity.ID) (*entity.PhieuNhap, error) {
	ctPhieuNhap, err := service.GetChiTietPhieuNhap(maChiTietPhieuNhap)

	if err != nil {
		return nil, err
	}

	phieuNhap, err := service.phieuNhapRepo.RemoveChiTietPhieuNhap(ctPhieuNhap)

	if err != nil {
		return nil, err
	}

	return phieuNhap, nil
}

func (service *NhapSachService) GetChiTietPhieuNhap(maChiTietPhieuNhap *entity.ID) (*entity.CtPhieuNhap, error) {
	ctPhieuNhap, err := service.phieuNhapRepo.GetChiTietPhieuNhap(maChiTietPhieuNhap)

	if err != nil {
		return nil, err
	}

	if ctPhieuNhap == nil {
		return nil, businessError.NewBusinessError("Chi tiet phieu nhap not found")
	}

	return ctPhieuNhap, nil
}
