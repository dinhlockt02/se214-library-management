package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"reflect"
)

type TheLoaiRepository struct {
	db *sqlx.DB
}

func NewTheLoaiRepository(db *sqlx.DB) *TheLoaiRepository {
	return &TheLoaiRepository{
		db: db,
	}
}

func (r *TheLoaiRepository) GetDanhSachTheLoai() (_ []*entity.TheLoai, err error) {
	tx := r.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	stmt, err := tx.Prepare(`SELECT MaTheLoai, TenTheLoai FROM TheLoai`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not prepare query", err)
	}
	rows, err := stmt.Query()
	defer func() {
		rows.Close()
	}()
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not execute query", err)
	}
	danhSachTheLoai := make([]*entity.TheLoai, 0)
	for rows.Next() {
		theLoai := &entity.TheLoai{}
		s := reflect.ValueOf(theLoai).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 1; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}
		maTheLoai := ""
		columns[0] = &maTheLoai
		err = rows.Scan(columns...)
		if err != nil {
			return nil, coreerror.NewInternalServerError("database error: scan rows failed", err)
		}
		theLoai.MaTheLoai, err = entity.StringToID(maTheLoai)
		danhSachTheLoai = append(danhSachTheLoai, theLoai)
	}
	return danhSachTheLoai, nil
}

func (r *TheLoaiRepository) GetTheLoai(maTheLoai *entity.ID) (_ *entity.TheLoai, err error) {
	tx := r.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt, err := tx.Prepare(`SELECT TenTheLoai FROM TheLoai WHERE MaTheLoai = ?`)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not prepare query", err)
	}
	row := stmt.QueryRow(maTheLoai.String())
	theLoai := &entity.TheLoai{}
	err = row.Scan(&(theLoai.TenTheLoai))
	if err == sql.ErrNoRows {
		return nil, coreerror.NewNotFoundError("the loai not found", err)
	} else if err != nil {
		return nil, coreerror.NewInternalServerError("database error: row scan failed", err)
	}
	theLoai.MaTheLoai = maTheLoai
	return theLoai, nil
}

func (r *TheLoaiRepository) CreateTheLoai(theLoai *entity.TheLoai) (_ *entity.TheLoai, err error) {
	tx := r.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	exec := `INSERT INTO TheLoai (MaTheLoai, TenTheLoai) VALUES (?, ?)`
	_, err = tx.Exec(exec, theLoai.MaTheLoai.String(), theLoai.TenTheLoai)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not insert into database", err)
	}
	return theLoai, nil
}

func (r *TheLoaiRepository) UpdateTheLoai(theLoai *entity.TheLoai) (_ *entity.TheLoai, err error) {
	tx := r.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	exec := `
	UPDATE TheLoai
	SET TenTheLoai = ?
	WHERE MaTheLoai = ?
	`
	_, err = tx.Exec(exec, theLoai.TenTheLoai, theLoai.MaTheLoai.String())
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not insert into database", err)
	}
	return theLoai, nil
}

func (r *TheLoaiRepository) RemoveTheLoai(maTheLoai *entity.ID) (err error) {
	tx := r.db.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	row := tx.QueryRow(`SELECT COUNT(*) FROM CT_TheLoai WHERE MaTheLoai = ?`, maTheLoai.String())
	var count int = 0
	err = row.Scan(&count)
	if err != nil {
		return coreerror.NewInternalServerError("database error: cant query number of rows", err)
	}
	if count > 0 {
		return coreerror.NewConflictError("can't not delete the loai because dau sach has this the loai", nil)
	}

	exec := `DELETE FROM TheLoai WHERE MaTheLoai = ?`
	_, err = tx.Exec(exec, maTheLoai.String())
	if err != nil {
		return coreerror.NewInternalServerError("database error: can't not insert into database", err)
	}
	return nil
}
