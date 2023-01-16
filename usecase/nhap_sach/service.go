package nhapsach

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
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

func (service *NhapSachService) CreatePhieuNhapSach(ngayLap *time.Time, ctPhieuNhapDto []dto.CtPhieuNhapDto) (*entity.PhieuNhap, error) {
	ctPhieuNhap := make([]*entity.CtPhieuNhap, len(ctPhieuNhapDto))
	var tongTien uint
	for i, ct := range ctPhieuNhapDto {
		id, err := entity.StringToID(ct.MaDauSach)
		if err != nil {
			return nil, err
		}
		dauSach, err := service.dauSachUsecase.GetDauSach(id)
		if err != nil {
			return nil, err
		}
		sach := entity.NewSach(dauSach, ct.NhaXuatBan, ct.TriGia, ct.NamXuatBan, ct.TinhTrang, ct.GhiChu)
		if !sach.IsValid() {
			return nil, coreerror.NewBadRequestError("Invalid sach data", nil)
		}
		ctPhieuNhap[i] = entity.NewCtPhieuNhap(sach, ct.DonGia)
		tongTien += ct.DonGia
	}

	phieuNhap := entity.NewPhieuNhap(ngayLap, tongTien, ctPhieuNhap)

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
	err := service.phieuNhapRepo.RemovePhieuNhap(maPhieuNhap)
	return err
}
