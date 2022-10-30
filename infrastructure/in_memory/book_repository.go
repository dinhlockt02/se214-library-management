package inmemory

import "daijoubuteam.xyz/se214-library-management/core/entity"

type InMemoryBookRepo struct {
	Sach []entity.Sach
}

func (db *InMemoryBookRepo) GetBooks() ([]*entity.Sach, error) {
	result := make([]*entity.Sach, len(db.Sach))
	for index := range db.Sach {
		result[index] = &db.Sach[index]
	}
	return result, nil
}

func (db *InMemoryBookRepo) CreateBook(book *entity.Sach) (*entity.Sach, error) {
	db.Sach = append(db.Sach, *book)
	return book, nil
}
