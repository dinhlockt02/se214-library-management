package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type DocGiaRepository interface {
	CreateDocGia(docGia *entity.DocGia) (*entity.DocGia, error)
	GetDocGia(maDocGia string) (*entity.DocGia, error)
	GetDanhSachDocGia() ([]*entity.DocGia, error)
	UpdateDocGia(docGia *entity.DocGia) (*entity.DocGia, error)
	RemoveDocGia(maDocGia string) error
}
