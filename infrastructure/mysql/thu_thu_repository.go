package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	"github.com/jmoiron/sqlx"
)

type ThuThuRepository struct {
	db *sqlx.DB
}

func NewThuThuRepository(db *sqlx.DB) *ThuThuRepository {
	return &ThuThuRepository{
		db: db,
	}
}

func (r *ThuThuRepository) GetDanhSachThuThu(query *repository.ThuThuSearchQuery) (_ []*entity.ThuThu, err error) {
	tx := r.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	stmt, err := tx.Prepare(`SELECT * FROM ThuThu WHERE Email LIKE '%?%' or PhoneNumber LIKE '%?%';`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not prepare query")
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not execute query")
	}
	defer rows.Close()

	danhSachThuThu := make([]*entity.ThuThu, 0)
	for rows.Next() {
		var thuThu entity.ThuThu
		err = rows.Scan(&thuThu)
		danhSachThuThu = append(danhSachThuThu, &thuThu)
	}
	if err != nil {
		return danhSachThuThu, coreerror.NewInternalServerError("database error: scan rows failed")
	}
	return danhSachThuThu, nil
}

func (r *ThuThuRepository) GetThuThu(maThuThu *entity.ID) (*entity.ThuThu, error) {
	panic("not implemented")
}
func (r *ThuThuRepository) CreateThuThu(thuThu *entity.ThuThu) (*entity.ThuThu, error) {
	panic("not implemented")
}
func (r *ThuThuRepository) UpdateThuThu(thuThu *entity.ThuThu) (*entity.ThuThu, error) {
	panic("not implemented")
}
