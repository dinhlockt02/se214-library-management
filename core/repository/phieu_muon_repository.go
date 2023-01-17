package repository

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
)

type PhieuMuonRepository interface {
	GetDanhSachPhieuMuon() ([]*entity.PhieuMuon, error)
	CreatePhieuMuon(phieuMuon *entity.PhieuMuon) (*entity.PhieuMuon, error)
}
