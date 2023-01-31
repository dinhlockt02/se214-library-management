package mysql

import (
	"database/sql"
	"fmt"
	"reflect"

	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
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

func (r *ThuThuRepository) GetDanhSachThuThu() (_ []*entity.ThuThu, err error) {

	tx := r.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt, err := tx.Prepare(
		`SELECT MaThuThu, Name, NgaySinh, Email, PhoneNumber, Password, Status, IsAdminRole
				FROM ThuThu`)

	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not prepare query", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not execute query", err)
	}
	defer rows.Close()

	danhSachThuThu := make([]*entity.ThuThu, 0)
	for rows.Next() {
		var maThuThu string = ""
		thuThu := &entity.ThuThu{}
		s := reflect.ValueOf(thuThu).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 1; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		columns[0] = &maThuThu
		err = rows.Scan(columns...)
		thuThu.MaThuThu, err = entity.StringToID(maThuThu)
		if err != nil {
			return danhSachThuThu, coreerror.NewInternalServerError("database error: scan rows failed", err)
		}
		danhSachThuThu = append(danhSachThuThu, thuThu)
	}
	return danhSachThuThu, nil
}
func (r *ThuThuRepository) GetThuThu(maThuThu *entity.ID) (_ *entity.ThuThu, err error) {
	tx := r.db.MustBegin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	stmt, err := tx.Prepare(`
		SELECT Name, NgaySinh, Email, PhoneNumber, Password, Status, IsAdminRole 
		FROM ThuThu 
		WHERE MaThuThu = ?
	`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: prepare query failed", err)
	}
	row := stmt.QueryRow(maThuThu.String())

	if err = row.Err(); err == sql.ErrNoRows {
		return nil, coreerror.NewNotFoundError("thu thu not found", err)
	} else if err != nil {
		return nil, coreerror.NewInternalServerError("database error: query failed", err)
	}

	thuThu := &entity.ThuThu{}
	err = row.Scan(
		&(thuThu.Name),
		&(thuThu.NgaySinh),
		&(thuThu.Email),
		&(thuThu.PhoneNumber),
		&(thuThu.Password),
		&(thuThu.Status),
		&(thuThu.IsAdminRole))
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, coreerror.NewNotFoundError("thu thu not found", err)
		}
		return nil, coreerror.NewInternalServerError("database error: can not scan row", err)
	}
	thuThu.MaThuThu = maThuThu
	return thuThu, nil
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
		return nil, coreerror.NewInternalServerError("database error: insert new thu thu failed", err)
	}

	return thuThu, nil
}
func (r *ThuThuRepository) UpdateThuThu(thuThu *entity.ThuThu) (*entity.ThuThu, error) {
	_, err := r.db.Exec(`
UPDATE ThuThu 
SET Password = ?,
    Name = ?,
    NgaySinh = ?,
    Email = ?,
    PhoneNumber = ?
WHERE MaThuThu = ?`, thuThu.Password, thuThu.Name, thuThu.NgaySinh, thuThu.Email, thuThu.PhoneNumber, thuThu.MaThuThu)
	return thuThu, err
}
func (r *ThuThuRepository) GetThuThuByEmail(email string) (_ *entity.ThuThu, err error) {
	tx := r.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt, err := tx.Prepare(`SELECT MaThuThu, Name, NgaySinh, Email, PhoneNumber, Password, Status, IsAdminRole 
									FROM ThuThu 
									WHERE Email LIKE ?`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: prepare query failed", err)
	}
	row := stmt.QueryRow(email)

	var maThuThu string = ""
	thuThu := &entity.ThuThu{}
	s := reflect.ValueOf(thuThu).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)
	for i := 1; i < numCols; i++ {
		field := s.Field(i)
		columns[i] = field.Addr().Interface()
	}
	columns[0] = &maThuThu
	err = row.Scan(columns...)
	if err == sql.ErrNoRows {
		return nil, coreerror.NewNotFoundError("thu thu not found", err)
	}
	thuThu.MaThuThu, err = entity.StringToID(maThuThu)
	if err != nil {
		return nil, coreerror.NewInternalServerError("databaser error: convert to id failed", err)
	}

	return thuThu, nil
}
