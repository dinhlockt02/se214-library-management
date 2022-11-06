package mysql

import (
	"fmt"
	"reflect"

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
	stmt, err := tx.Preparex(
		`SELECT MaThuThu, Name, NgaySinh, Email, PhoneNumber, Password, Status, IsAdminRole FROM ThuThu`,
	)

	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not prepare query")
	}
	rows, err := stmt.Queryx()
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not execute query")
	}
	defer rows.Close()

	danhSachThuThu := make([]*entity.ThuThu, 0)
	for rows.Next() {
		var maThuThu string = ""
		thuthu := &entity.ThuThu{}
		s := reflect.ValueOf(thuthu).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 1; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		columns[0] = &maThuThu
		err = rows.Scan(columns...)
		thuthu.MaThuThu, err = entity.StringToID(maThuThu)
		fmt.Println(err)
		danhSachThuThu = append(danhSachThuThu, thuthu)
	}

	if err != nil {
		return danhSachThuThu, coreerror.NewInternalServerError("database error: scan rows failed")
	}
	return danhSachThuThu, nil
}

func (r *ThuThuRepository) GetThuThu(maThuThu *entity.ID) (*entity.ThuThu, error) {
	panic("not implemented")
}
func (r *ThuThuRepository) CreateThuThu(thuThu *entity.ThuThu) (_ *entity.ThuThu, err error) {
	tx := r.db.MustBegin()

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	thuThuExec := `INSERT INTO ThuThu (MaThuThu, Name, NgaySinh, Email, PhoneNumber, Password, Status, IsAdminRole) VALUES (?, ? ,?, ?, ? , ? , ? , ?)`

	_, err = tx.Exec(thuThuExec, thuThu.MaThuThu, thuThu.Name, thuThu.NgaySinh, thuThu.Email, thuThu.PhoneNumber, thuThu.Password, thuThu.Status, thuThu.IsAdminRole)

	if err != nil {
		return nil, err
	}

	return thuThu, nil
}
func (r *ThuThuRepository) UpdateThuThu(thuThu *entity.ThuThu) (*entity.ThuThu, error) {
	panic("not implemented")
}
