package tra_sach

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	"daijoubuteam.xyz/se214-library-management/usecase/muon_sach"
	"time"
)

type Service struct {
	repo            repository.PhieuTraRepository
	muonSachUsecase muon_sach.Usecase
}

func (s *Service) GetPhieuTra() ([]*entity.PhieuTra, error) {
	return s.repo.GetDanhSachPhieuTra()
}

func (s *Service) CreatePhieuTra(maSach *entity.ID, ghiChu string, ngayTra *time.Time) (*entity.PhieuTra, error) {
	phieuMuon, err := s.muonSachUsecase.GetPhieuMuonByMaSach(maSach)
	if err != nil {
		return nil, err
	}
	diff := diffDay(phieuMuon.NgayMuon, ngayTra)
	if diff < 0 {
		return nil, coreerror.NewBadRequestError("Invalid ngay tra", nil)
	}
	var tienPhat uint
	if uint(diff) > phieuMuon.LoaiDocGia.ThoiGianMuonToiDa {
		tienPhat = (uint(diff) - phieuMuon.LoaiDocGia.ThoiGianMuonToiDa) * phieuMuon.LoaiDocGia.TienPhatTheoNgay
	}
	phieuMuon.DocGia.TongNo += int(tienPhat)
	return s.repo.CreatePhieuTra(&entity.PhieuTra{
		TienPhat:  tienPhat,
		NgayTra:   ngayTra,
		GhiChu:    ghiChu,
		PhieuMuon: phieuMuon,
	})
}

func diffDay(d1 *time.Time, d2 *time.Time) int {
	duration := d2.Sub(*d1)
	return int(duration.Hours() / 24)
}

func NewTraSachService(repo repository.PhieuTraRepository, muonSachUsecase muon_sach.Usecase) *Service {
	return &Service{
		repo:            repo,
		muonSachUsecase: muonSachUsecase,
	}
}
