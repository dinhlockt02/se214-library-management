package muon_sach

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"time"
)

type Usecase interface {
	GetPhieuMuon() ([]*entity.PhieuMuon, error)
	GetPhieuMuonByDocGia(maDocGia string) ([]*entity.PhieuMuon, error)
	GetPhieuMuonByMaSach(maSach *entity.ID) (*entity.PhieuMuon, error)
	CreatePhieuMuon(ngayMuon *time.Time, maSach *entity.ID, maDocGia string) (*entity.PhieuMuon, error)
}
