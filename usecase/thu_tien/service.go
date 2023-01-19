package thu_tien

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	docgia "daijoubuteam.xyz/se214-library-management/usecase/doc_gia"
	"time"
)

type Service struct {
	repo          repository.PhieuThuTienRepository
	docGiaUsecase docgia.DocGiaUsecase
}

func NewThuTienService(repo repository.PhieuThuTienRepository, docGiaUsecase docgia.DocGiaUsecase) Service {
	return Service{
		repo:          repo,
		docGiaUsecase: docGiaUsecase,
	}
}

func (s Service) GetPhieuThuTienByMaDocGia(maDocGia string) (_ []*entity.PhieuThuTien, err error) {
	return s.repo.GetPhieuThuTienByMaDocGia(maDocGia)
}

func (s Service) CreatePhieuThuTien(maDocGia string, soTienThu int, ngayThu *time.Time) (_ *entity.PhieuThuTien, err error) {
	docGia, err := s.docGiaUsecase.GetDocGia(maDocGia)
	if err != nil {
		return nil, err
	}
	tongNo := docGia.TongNo
	conLai := docGia.TongNo - soTienThu
	phieuThu := entity.NewPhieuThuTien(docGia, ngayThu, tongNo, soTienThu, conLai)
	docGia.TongNo = conLai
	return s.repo.CreatePhieuThuTien(phieuThu)
}
