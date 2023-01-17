package muon_sach

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	docgia "daijoubuteam.xyz/se214-library-management/usecase/doc_gia"
	"daijoubuteam.xyz/se214-library-management/usecase/sach"
	"time"
)

type Service struct {
	SachUsecase         sach.SachUsecase
	DocGiaUsecase       docgia.DocGiaUsecase
	PhieuMuonRepository repository.PhieuMuonRepository
}

func NewMuonSachService(sachUsecase sach.SachUsecase,
	phieuMuonRepository repository.PhieuMuonRepository,
	docGiaUsecase docgia.DocGiaUsecase) *Service {
	return &Service{
		SachUsecase:         sachUsecase,
		PhieuMuonRepository: phieuMuonRepository,
		DocGiaUsecase:       docGiaUsecase,
	}
}

func (s Service) GetPhieuMuon() ([]*entity.PhieuMuon, error) {
	return s.PhieuMuonRepository.GetDanhSachPhieuMuon()
}

func (s Service) CreatePhieuMuon(ngayMuon *time.Time, maSach *entity.ID, maDocGia *entity.ID) (*entity.PhieuMuon, error) {
	var err error
	var dg *entity.DocGia
	if dg, err = s.DocGiaUsecase.GetDocGia(maDocGia); err != nil {
		return nil, err
	}
	var sa *entity.Sach
	if sa, err = s.SachUsecase.GetSach(maSach); err != nil {
		return nil, err
	}
	if !sa.TinhTrang {
		return nil, coreerror.NewConflictError("conflict error: sach da duoc muon", nil)
	}
	sa.TinhTrang = false
	return s.PhieuMuonRepository.CreatePhieuMuon(entity.NewPhieuMuon(dg, ngayMuon, sa, nil))
}
