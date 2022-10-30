package book

import "daijoubuteam.xyz/se214-library-management/core/entity"

type BookUsecase interface {
	GetBooks() ([]*entity.Sach, error)
	CreateBook(sach *entity.Sach) (*entity.Sach, error)
}
